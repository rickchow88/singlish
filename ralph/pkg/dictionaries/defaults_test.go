package dictionaries

import (
	"testing"
)

func TestGetDefaultMappings(t *testing.T) {
	defaults := GetDefaultMappings()

	if defaults == nil {
		t.Fatal("GetDefaultMappings() returned nil")
	}

	expectedCount := 18
	if len(defaults) != expectedCount {
		t.Errorf("GetDefaultMappings() returned %d entries, want %d", len(defaults), expectedCount)
	}

	tests := []struct {
		singlish string
		wantGo   string
	}{
		{"kampung", "package"},
		{"action", "func"},
		{"gong", "fmt.Println"},
		{"nasi", "if"},
		{"dapao", "import"},
		{"bolehtak", "bool"},
		{"nombor", "int"},
		{"cheem", "complex128"},
		{"zhi", "rune"},
		{"tar", "string"},
		{"ki", "*"},
		{"buat", "make"},
		{"kwear", "close"},
		{"buang", "delete"},
		{"upsize", "append"},
		{"pattern", "type"},
		{"susun", "struct"},
		{"lobang", "interface"},
	}

	for _, tt := range tests {
		if got, ok := defaults[tt.singlish]; !ok || got != tt.wantGo {
			t.Errorf("GetDefaultMappings()[%q] = %q, want %q", tt.singlish, got, tt.wantGo)
		}
	}
}

func TestNewDefaultDictionary(t *testing.T) {
	dict := NewDefaultDictionary()
	if dict == nil {
		t.Fatal("NewDefaultDictionary() returned nil")
	}

	// Verify counts
	defaults := GetDefaultMappings()
	if len(dict.mapping) != len(defaults) {
		t.Errorf("NewDefaultDictionary() mapping count = %d, want %d", len(dict.mapping), len(defaults))
	}

	// Verify a few mappings
	if got, found := dict.Lookup("kampung"); !found || got != "package" {
		t.Errorf("Lookup('kampung') = %q, want 'package'", got)
	}

	// Verify reverse mapping logic
	if got, found := dict.ReverseLookup("package"); !found || got != "kampung" {
		t.Errorf("ReverseLookup('package') = %q, want 'kampung'", got)
	}
}
