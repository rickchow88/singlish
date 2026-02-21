package formatter

import (
	"os"
	"testing"

	"github.com/rickchow/singlish/pkg/ast"
	"github.com/rickchow/singlish/pkg/dictionaries"
	"github.com/rickchow/singlish/pkg/lexer"
)

func TestFormat(t *testing.T) {
	// 1. Create a dummy dictionary
	dictContent := `
kampung: package
import: import
action: func
dunia: var
siam: return
nombor: int
`
	tmpDir := t.TempDir()
	dictPath := tmpDir + "/SINGLISH_KEYWORDS.md"
	if err := os.WriteFile(dictPath, []byte(dictContent), 0644); err != nil {
		t.Fatalf("Failed to create temp dict: %v", err)
	}

	dict, err := dictionaries.LoadDictionary(dictPath)
	if err != nil {
		t.Fatalf("Failed to load dictionary: %v", err)
	}

	// 2. Construct AST manually or assume Parser works
	// Let's assume we can construct a simple AST
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.PackageStatement{
				Token: lexer.Token{Type: lexer.TokenKeyword, Value: "package"}, // Use Go keyword
				Name:  &ast.Identifier{Value: "main"},
			},
			&ast.FunctionStatement{
				Token: lexer.Token{Type: lexer.TokenKeyword, Value: "func"}, // Use Go keyword
				Name:  &ast.Identifier{Value: "main"},
				Body: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.LetStatement{
							Token: lexer.Token{Type: lexer.TokenKeyword, Value: "var"}, // Use Go keyword
							Names: []*ast.Identifier{{Value: "x"}},
							Type:  &ast.Identifier{Value: "int"}, // Use Go type
							Value: &ast.IntegerLiteral{Token: lexer.Token{Type: lexer.TokenNumber, Value: "1"}, Value: 1},
						},
						&ast.ReturnStatement{
							Token: lexer.Token{Type: lexer.TokenKeyword, Value: "return"}, // Use Go keyword
						},
					},
				},
			},
		},
	}

	// 3. Format
	output, err := Format(program, dict)
	if err != nil {
		t.Fatalf("Format failed: %v", err)
	}

	// 4. Verify canonicalization
	expected := `kampung main

action main() {
	dunia x nombor = 1
	siam
}
`
	// Note: indentation is tabs.
	// Check package -> kampung
	// Check func -> action
	// Check var -> dunia
	// Check int -> nombor (if implemented)
	// Check return -> siam
	// Check indentation

	if output != expected {
		t.Errorf("Format output mismatch.\nExpected:\n%q\nGot:\n%q", expected, output)
	}
}
