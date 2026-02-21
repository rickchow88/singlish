package reporting

import (
	"fmt"
	"io"
	"strings"

	"github.com/rickchow/singlish/pkg/lexer"
)

// PrintErrorWithContext prints a diagnostic with the source line and a caret pointing to the error.
func PrintErrorWithContext(out io.Writer, source string, diag lexer.Diagnostic) {
	lines := strings.Split(source, "\n")
	lineIndex := diag.Line - 1 // Line is 1-based

	fmt.Fprintf(out, "Error on line %d: %s\n", diag.Line, diag.Message)

	if lineIndex >= 0 && lineIndex < len(lines) {
		line := lines[lineIndex]

		// Print the line code
		fmt.Fprintf(out, "%s\n", line)

		// Create pointer string
		// diag.Col is 1-based.
		pad := ""
		if diag.Col > 1 {
			// We need to replicate the indentation (tabs/spaces) to ensure caret aligns correctly.
			// Taking the substring up to the error column allows us to mirror whitespace.
			if diag.Col-1 <= len(line) {
				prefix := line[:diag.Col-1]
				for _, r := range prefix {
					if r == '\t' {
						pad += "\t"
					} else {
						pad += " "
					}
				}
			} else {
				// Fallback if column is out of bounds (shouldn't happen ideally)
				pad = strings.Repeat(" ", diag.Col-1)
			}
		}

		marker := "^"
		if diag.Length > 1 {
			marker = strings.Repeat("^", diag.Length)
		}

		fmt.Fprintf(out, "%s%s\n", pad, marker)
	}
}

// PrintDiagnostics prints multiple diagnostics.
func PrintDiagnostics(out io.Writer, source string, diags []lexer.Diagnostic) {
	for _, d := range diags {
		PrintErrorWithContext(out, source, d)
	}
}
