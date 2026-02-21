package integration

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestChaosErrors(t *testing.T) {
	// Create a temp file with invalid Singlish code
	tempFile, err := os.CreateTemp("", "chaos_*.singlish")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// "kopi" is not a valid keyword (unless it was added recently, let's assume valid keywords are in docs)
	// Actually, let's use a completely random string to ensure it's invalid
	// Or a syntax error like an unclosed string
	invalidCode := `
kampung chaos

action main() {
  gong("This string is not closed
}
`
	if _, err := tempFile.WriteString(invalidCode); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Run singlish run on the invalid file
	cmd := exec.Command(singlishBinary, "run", tempFile.Name())
	// Set working directory to the binary's directory (which is clean/empty)
	// to ensure we are not relying on any dictionary file in the repo root.
	cmd.Dir = filepath.Dir(singlishBinary)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	// We expect failure
	err = cmd.Run()
	if err == nil {
		t.Error("Expected command to fail on invalid input, but it succeeded")
	}

	output := stderr.String()
	t.Logf("Stderr output:\n%s", output)

	// Check for insult
	foundInsult := false
	insults := []string{
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

	for _, insult := range insults {
		if strings.Contains(output, insult) {
			foundInsult = true
			break
		}
	}

	if !foundInsult {
		t.Error("Expected output to contain a Singlish insult, but none found")
	}

	// Check for original error text
	// The transpiler (lexer) should report unterminated string
	// Or if I used a random token, it might report unexpected token
	if !strings.Contains(output, "Error:") && !strings.Contains(output, "error") {
		t.Error("Expected output to contain original error message")
	}
}
