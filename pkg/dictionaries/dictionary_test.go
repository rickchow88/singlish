package dictionaries

import (
	"os"
	"path/filepath"
	"testing"
)

// Helper to create a temporary dictionary file for testing
func createTempDictionaryFile(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	tmpFilePath := filepath.Join(tmpDir, "SINGLISH_KEYWORDS.md")
	err := os.WriteFile(tmpFilePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write temporary dictionary file: %v", err)
	}
	return tmpFilePath
}

func TestLoadDictionary(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
		wantMap map[string]string
	}{
		{
			name: "valid dictionary",
			content: `
kampung: package
action: func
// A comment
nasi: if
`,
			wantErr: false,
			wantMap: map[string]string{
				"kampung": "package",
				"action":  "func",
				"nasi":    "if",
			},
		},
		{
			name: "empty dictionary",
			content: `
// Just comments
`,
			wantErr: false,
			wantMap: map[string]string{},
		},
		{
			name:    "invalid entry format",
			content: `kampung package`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFilePath := createTempDictionaryFile(t, tt.content)

			dict, err := LoadDictionary(tmpFilePath)

			if (err != nil) != tt.wantErr {
				t.Errorf("LoadDictionary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if dict == nil {
					t.Fatal("LoadDictionary() returned nil dictionary for no error")
				}
				if len(dict.mapping) != len(tt.wantMap) {
					t.Errorf("LoadDictionary() got %d entries, want %d", len(dict.mapping), len(tt.wantMap))
				}
				for k, v := range tt.wantMap {
					if got, ok := dict.mapping[k]; !ok || got != v {
						t.Errorf("LoadDictionary() mapping[%q] = %q, want %q", k, got, v)
					}
				}
			}
		})
	}
}

func TestLookup(t *testing.T) {
	// Create a dummy dictionary for lookup tests
	content := `
kampung: package
action: func
nasi: if
`
	tmpFilePath := createTempDictionaryFile(t, content)

	dict, err := LoadDictionary(tmpFilePath)
	if err != nil {
		t.Fatalf("Failed to load dictionary for Lookup tests: %v", err)
	}

	tests := []struct {
		name          string
		singlish      string
		wantGoKeyword string
		wantFound     bool
	}{
		{
			name:          "existing keyword",
			singlish:      "kampung",
			wantGoKeyword: "package",
			wantFound:     true,
		},
		{
			name:          "another existing keyword",
			singlish:      "action",
			wantGoKeyword: "func",
			wantFound:     true,
		},
		{
			name:          "non-existing keyword",
			singlish:      "kampungers",
			wantGoKeyword: "",
			wantFound:     false,
		},
		{
			name:          "empty string lookup",
			singlish:      "",
			wantGoKeyword: "",
			wantFound:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGoKeyword, gotFound := dict.Lookup(tt.singlish)
			if gotGoKeyword != tt.wantGoKeyword {
				t.Errorf("Lookup(%q) got Go keyword %q, want %q", tt.singlish, gotGoKeyword, tt.wantGoKeyword)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Lookup(%q) got found %v, want %v", tt.singlish, gotFound, tt.wantFound)
			}
		})
	}
}
