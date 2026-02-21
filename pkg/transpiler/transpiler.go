package transpiler

import (
	"fmt"

	"github.com/rickchow/singlish/pkg/codegen"
	"github.com/rickchow/singlish/pkg/dictionaries"
	"github.com/rickchow/singlish/pkg/lexer"
	"github.com/rickchow/singlish/pkg/parser"
)

// TranspilationError wraps a list of diagnostics from lexer or parser.
type TranspilationError struct {
	Diagnostics []lexer.Diagnostic
}

func (e *TranspilationError) Error() string {
	if len(e.Diagnostics) == 0 {
		return "transpilation failed"
	}
	// Return the first error message as the default string representation.
	return fmt.Sprintf("%s (and %d more errors)", e.Diagnostics[0].Message, len(e.Diagnostics)-1)
}

// Transpile converts Singlish source code to Go source code.
// It uses the AST-based pipeline: Lexer -> Parser -> Codegen.
func Transpile(source string, dict *dictionaries.Dictionary) (string, error) {
	// 1. Lex
	keywords := make(map[string]struct{})
	if dict != nil {
		for _, k := range dict.Keys() {
			keywords[k] = struct{}{}
		}
	}
	// Ensure 'ki' is treated as a keyword for pointer syntax
	keywords["ki"] = struct{}{}

	tokens, diagnostics := lexer.Lex(source, keywords)
	if len(diagnostics) > 0 {
		return "", &TranspilationError{Diagnostics: diagnostics}
	}

	// 2. Parse
	p := parser.New(tokens, dict)
	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		return "", &TranspilationError{Diagnostics: p.Errors()}
	}

	// 3. Codegen
	code, err := codegen.Generate(program, dict)
	if err != nil {
		return "", fmt.Errorf("codegen error: %w", err)
	}

	return code, nil
}
