package cmd

import (
	"fmt"
	"os"

	"github.com/rickchow/singlish/pkg/formatter"
	"github.com/rickchow/singlish/pkg/lexer"
	"github.com/rickchow/singlish/pkg/parser"
)

const fmtUsage = `Usage:
  singlish fmt <file>

Description:
  Format the Singlish source file using canonical Singlish keywords and standard indentation.
`

func runFmt(args []string) int {
	if len(args) == 0 || isHelpFlag(args[0]) {
		fmt.Fprint(os.Stdout, fmtUsage)
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "\nError: missing input file")
			return 1
		}
		return 0
	}

	inputFile := args[0]
	if err := formatFile(inputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	return 0
}

func formatFile(inputPath string) error {
	// Read input
	content, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Load dictionary
	dict, err := loadDictionary()
	if err != nil {
		return fmt.Errorf("failed to load dictionary: %w", err)
	}

	// Lex
	keywords := make(map[string]struct{})
	for _, k := range dict.Keys() {
		keywords[k] = struct{}{}
	}
	// Ensure 'ki' is treated as a keyword
	keywords["ki"] = struct{}{}

	tokens, diagnostics := lexer.Lex(string(content), keywords)
	if len(diagnostics) > 0 {
		return fmt.Errorf("lexer error: %v", diagnostics[0].Message)
	}

	// Parse
	p := parser.New(tokens, dict)
	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		return fmt.Errorf("parser error: %v", p.Errors()[0])
	}

	// Format
	formatted, err := formatter.Format(program, dict)
	if err != nil {
		return fmt.Errorf("formatting failed: %w", err)
	}

	// Write back to file
	if err := os.WriteFile(inputPath, []byte(formatted), 0644); err != nil {
		return fmt.Errorf("failed to write formatted file: %w", err)
	}

	return nil
}
