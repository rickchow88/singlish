package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"

	"github.com/rickchow/singlish/pkg/dictionaries"
	"github.com/rickchow/singlish/pkg/reporting"
	"github.com/rickchow/singlish/pkg/transpiler"
)

var insults = []string{
	"Eh bodoh",
	"Go fly kite lah",
	"Simi sai is this?",
	"You think this one playground ah?",
	"Catch no ball sia",
	"Blur like sotong",
	"Why you so liddat?",
	"Aiyo, cannot make it lah",
	"Wake up your idea",
}

func getRandomInsult() string {
	return insults[rand.Intn(len(insults))]
}

// executeWithInsults runs the command and wraps stderr with an insult if it fails.
func executeWithInsults(cmd *exec.Cmd) error {
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		// Command failed, print insult then error
		fmt.Fprintf(os.Stderr, "%s\n", getRandomInsult())
		fmt.Fprint(os.Stderr, stderr.String())
		return err
	}

	// Command succeeded, just print any stderr output (e.g. warnings)
	if stderr.Len() > 0 {
		fmt.Fprint(os.Stderr, stderr.String())
	}
	return nil
}

// printErrorWithInsult prints a random insult followed by the error message.
func printErrorWithInsult(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", getRandomInsult())
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}

// handleError prints rich diagnostics if available, or falls back to insults.
func handleError(err error, inputPath string) {
	var tErr *transpiler.TranspilationError
	if errors.As(err, &tErr) {
		content, readErr := os.ReadFile(inputPath)
		if readErr == nil {
			reporting.PrintDiagnostics(os.Stderr, string(content), tErr.Diagnostics)
			return
		}
	}
	printErrorWithInsult(err)
}

func loadDictionary() (*dictionaries.Dictionary, error) {
	// 1. Explicit flag
	if DictionaryPath != "" {
		return dictionaries.LoadDictionary(DictionaryPath)
	}

	// 2. Try environment variable
	if path := os.Getenv("SINGLISH_KEYWORDS"); path != "" {
		return dictionaries.LoadDictionary(path)
	}

	// 3. Use embedded default
	return dictionaries.NewDefaultDictionary(), nil
}

// transpileToTemp reads inputPath, transpiles it, and writes the result to a temporary .go file.
// It returns the path to the temp file and any error.
func transpileToTemp(inputPath string) (string, error) {
	// Read input
	content, err := os.ReadFile(inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to read input file: %w", err)
	}

	// Load dictionary
	dict, err := loadDictionary()
	if err != nil {
		return "", fmt.Errorf("failed to load dictionary: %w", err)
	}

	// Transpile
	goCode, err := transpiler.Transpile(string(content), dict)
	if err != nil {
		return "", fmt.Errorf("transpilation failed: %w", err)
	}

	// Write to temp file
	tmpFile, err := os.CreateTemp("", "singlish_*.go")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	defer tmpFile.Close()

	if _, err := tmpFile.WriteString(goCode); err != nil {
		return "", fmt.Errorf("failed to write to temp file: %w", err)
	}

	return tmpFile.Name(), nil
}
