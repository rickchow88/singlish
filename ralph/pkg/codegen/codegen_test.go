package codegen

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rickchow/singlish/pkg/dictionaries"
	"github.com/rickchow/singlish/pkg/lexer"
	"github.com/rickchow/singlish/pkg/parser"
)

func createTempDictionaryFile(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	tmpFilePath := filepath.Join(tmpDir, "SINGLISH_KEYWORDS.md")
	err := os.WriteFile(tmpFilePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write temporary dictionary file: %v", err)
	}
	return tmpFilePath
}

func TestGenerate(t *testing.T) {
	dictContent := `
kampung: package
importlah: import
dun_var: var
give_back: return
`
	dictPath := createTempDictionaryFile(t, dictContent)
	dict, err := dictionaries.LoadDictionary(dictPath)
	if err != nil {
		t.Fatalf("Failed to load dictionary: %v", err)
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Basic package and var",
			input: `
kampung main
dun_var x = 5;
`,
			expected: `package main

var x = 5
`,
		},
		{
			name: "Auto-import fmt",
			input: `
kampung main
fmt.Println("Hello");
`,
			expected: `package main

import (
	"fmt"
)

fmt.Println("Hello")
`,
		},
		{
			name: "Explicit import",
			input: `
kampung main
importlah "fmt"
fmt.Println("Hello");
`,
			expected: `package main

import (
	"fmt"
)

fmt.Println("Hello")
`,
		},
		{
			name: "Precedence",
			input: `
kampung main
dun_var x = 1 + 2 * 3;
dun_var y = (1 + 2) * 3;
`,
			expected: `package main

var x = (1 + (2 * 3))
var y = ((1 + 2) * 3)
`,
		},
		{
			name: "Return statement",
			input: `
give_back 5;
`,
			expected: `package main

return 5
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Lexer needs keywords map
			keywords := make(map[string]struct{})
			for _, k := range dict.Keys() {
				keywords[k] = struct{}{}
			}

			tokens, diags := lexer.Lex(tt.input, keywords)
			if len(diags) > 0 {
				t.Fatalf("Lexer error: %v", diags)
			}

			p := parser.New(tokens, dict)
			program := p.ParseProgram()
			if len(p.Errors()) > 0 {
				t.Fatalf("Parser errors: %v", p.Errors())
			}

			got, err := Generate(program, dict)
			if err != nil {
				t.Fatalf("Generate error: %v", err)
			}

			if got != tt.expected {
				t.Errorf("Generate() mismatch.\nWant:\n%q\nGot:\n%q", tt.expected, got)
			}
		})
	}
}
