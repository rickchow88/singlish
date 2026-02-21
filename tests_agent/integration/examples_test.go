package integration

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var singlishBinary string

func TestMain(m *testing.M) {
	// Setup: Build the singlish binary
	tempDir, err := os.MkdirTemp("", "singlish-test-build")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create temp dir: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tempDir)

	exeName := "singlish"
	if runtime.GOOS == "windows" {
		exeName += ".exe"
	}
	singlishBinary = filepath.Join(tempDir, exeName)

	// Determine repo root from current test file location
	// We assume test is running in tests/integration
	repoRoot, err := filepath.Abs("../../")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to resolve repo root: %v\n", err)
		os.Exit(1)
	}

	mainPath := filepath.Join(repoRoot, "main.go")

	// Build the CLI
	cmd := exec.Command("go", "build", "-o", singlishBinary, mainPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to build singlish binary: %v\nOutput:\n%s\n", err, output)
		os.Exit(1)
	}

	// Run tests
	code := m.Run()

	os.Exit(code)
}

func TestExamples(t *testing.T) {
	// Point to the root docs/examples directory (3 levels up from ralph/tests/integration)
	examplesDir, err := filepath.Abs("../../../docs/examples")
	if err != nil {
		t.Fatalf("Failed to resolve examples directory: %v", err)
	}

	// files to skip
	skipFiles := map[string]bool{
		"debug_catch.singlish": true, // Intentionally blocks/deadlocks
	}

	// specific output expectations (partial match)
	expectedOutputs := map[string]string{
		"01_hello_world.singlish": "Hello Singapore! Limpeh is coding now.",
		"02_fizzbuzz.singlish":    "FizzBuzz",
		"11_pointers.singlish":    "60",
	}

	files, err := filepath.Glob(filepath.Join(examplesDir, "*.singlish"))
	if err != nil {
		t.Fatalf("Failed to list example files: %v", err)
	}

	if len(files) == 0 {
		t.Fatal("No example files found")
	}

	for _, path := range files {
		filename := filepath.Base(path)
		if skipFiles[filename] {
			t.Logf("Skipping %s", filename)
			continue
		}

		t.Run(filename, func(t *testing.T) {
			// Cmd to run
			cmd := exec.Command(singlishBinary, "run", path)
			cmd.Dir = filepath.Dir(singlishBinary) // Run in clean dir

			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr

			err := cmd.Run()
			if err != nil {
				t.Fatalf("Command failed: %v\nStderr: %s\nStdout: %s", err, stderr.String(), out.String())
			}

			output := out.String()

			// Check against expected output if defined
			if expected, ok := expectedOutputs[filename]; ok {
				if !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain %q, got:\n%s", expected, output)
				}
			}
		})
	}
}

func TestMissingFile(t *testing.T) {
	cmd := exec.Command(singlishBinary, "run", "missing.singlish")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err == nil {
		t.Error("Expected error for missing file, got success")
	}

	errOut := stderr.String()
	if errOut == "" {
		t.Error("Expected error message on stderr, got empty string")
	}
}
