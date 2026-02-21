package transpiler

import (
	"strings"
	"testing"
)

func TestTranspileStructAndInterface(t *testing.T) {
	dict := createTempDictionary(t, `
kampung: package
pattern: type
susun: struct
lobang: interface
nombor: int
tar: string
`)

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "Simple Struct",
			input: `
kampung main
pattern User susun {
	Name tar
	Age nombor
}
`,
			expected: []string{
				"type User struct {",
				"Name string",
				"Age int",
				"}",
			},
		},
		{
			name: "Struct with Tags",
			input: `
kampung main
pattern Config susun {
	Host tar "json:\"host\""
	Port nombor "json:\"port\""
}
`,
			expected: []string{
				"type Config struct {",
				"Host string \"json:\\\"host\\\"\"",
				"Port int \"json:\\\"port\\\"\"",
				"}",
			},
		},
		{
			name: "Simple Interface",
			input: `
kampung main
pattern Speaker lobang {
	Speak() tar
	Listen(msg tar)
}
`,
			expected: []string{
				"type Speaker interface {",
				"Speak() string",
				"Listen(msg string)",
				"}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Transpile(tt.input, dict)
			if err != nil {
				t.Fatalf("Transpile failed: %v", err)
			}

			for _, exp := range tt.expected {
				if !strings.Contains(got, exp) {
					t.Errorf("Expected output to contain %q, but got:\n%s", exp, got)
				}
			}
		})
	}
}
