package transpiler

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rickchow/singlish/pkg/dictionaries"
)

func createTempDictionary(t *testing.T, content string) *dictionaries.Dictionary {
	t.Helper()
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "SINGLISH_KEYWORDS.md")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("Failed to write dict file: %v", err)
	}
	dict, err := dictionaries.LoadDictionary(path)
	if err != nil {
		t.Fatalf("Failed to load dict: %v", err)
	}
	return dict
}

func TestTranspileBasic(t *testing.T) {
	dict := createTempDictionary(t, "kampung: package\ngong: fmt.Println\n")

	input := "kampung main\n\ngong(\"hello\")"
	expectedContains := []string{
		"package main",
		"fmt.Println(\"hello\")",
		"\"fmt\"", // Just check for the package name in string
	}

	got, err := Transpile(input, dict)
	if err != nil {
		t.Fatalf("Transpile failed: %v", err)
	}

	if !strings.Contains(got, "import") {
		t.Error("Missing import statement")
	}

	for _, exp := range expectedContains {
		if !strings.Contains(got, exp) {
			t.Errorf("Expected output to contain %q, but got:\n%s", exp, got)
		}
	}
}

func TestTranspileContent(t *testing.T) {
	dict := createTempDictionary(t, "kampung: package\n")

	// Input with comments and whitespace
	input := "kampung main\n\n// comment\nfunc foo() {}"

	got, err := Transpile(input, dict)
	if err != nil {
		t.Fatalf("Transpile failed: %v", err)
	}

	// Verify key components are present
	if !strings.Contains(got, "package main") {
		t.Error("Missing package main")
	}
	// 'func' and 'foo()' might be on separate lines if parsed as separate expressions
	if !strings.Contains(got, "func") {
		t.Error("Missing func")
	}
	if !strings.Contains(got, "foo()") {
		t.Error("Missing foo()")
	}
	// Comments are currently skipped by Parser, so we don't check for "// comment"
}

func TestTranspileNoInjectFmtIfUnused(t *testing.T) {
	dict := createTempDictionary(t, "kampung: package\n")

	input := "kampung main"
	got, err := Transpile(input, dict)
	if err != nil {
		t.Fatalf("Transpile failed: %v", err)
	}

	if strings.Contains(got, "import \"fmt\"") {
		t.Errorf("Should not inject fmt if not used. Got:\n%s", got)
	}
}

func TestTranspileNoInjectFmtIfPresent(t *testing.T) {
	dict := createTempDictionary(t, "kampung: package\ngong: fmt.Println\ndapao: import\n")

	// existing import
	// Use extra space to ensure separation even if keyword expands (dapao=5, import=6)
	input := "kampung main\ndapao  \"fmt\"\ngong(\"hi\")"

	got, err := Transpile(input, dict)
	if err != nil {
		t.Fatalf("Transpile failed: %v", err)
	}

	// We expect fmt.Println
	if !strings.Contains(got, "fmt.Println") {
		t.Error("Missing fmt.Println")
	}

	// We expect "import" mapped from "dapao"
	// Output format is import (\n\t"fmt"\n)
	if !strings.Contains(got, "\"fmt\"") {
		t.Errorf("Missing \"fmt\" in output, got:\n%s", got)
	}
	// Check that we have an import block
	if !strings.Contains(got, "import (") && !strings.Contains(got, "import \"fmt\"") {
		t.Errorf("Missing import statement, got:\n%s", got)
	}

	// We check for duplication (count occurences of "fmt")
	// "import" "fmt" -> 1
	// "fmt.Println" -> 1
	// Total 2 "fmt" strings (roughly).

	lines := strings.Split(got, "\n")
	if strings.Contains(lines[0], "import") {
		t.Errorf("Should not have injected import into package line (line 1). Got: %q", lines[0])
	}
}

func TestTranspilePointer(t *testing.T) {
	// "tahan" -> "var"
	// "nombor" -> "int"
	// "ki" -> pointer
	dict := createTempDictionary(t, "tahan: var\nnombor: int\n")

	// tahan x ki nombor
	input := "kampung main\ntahan x ki nombor"

	got, err := Transpile(input, dict)
	if err != nil {
		t.Fatalf("Transpile failed: %v", err)
	}

	if !strings.Contains(got, "var x *int") {
		t.Errorf("Expected 'var x *int', got:\n%s", got)
	}
}
