package dictionaries

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Dictionary maps Singlish keywords to Go keywords.
type Dictionary struct {
	mapping        map[string]string
	reverseMapping map[string]string
}

// LoadDictionary reads the dictionary file from the given filePath and returns a new Dictionary.
func LoadDictionary(filePath string) (*Dictionary, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open dictionary file %q: %w", filePath, err)
	}
	defer file.Close()

	dict := &Dictionary{
		mapping:        make(map[string]string),
		reverseMapping: make(map[string]string),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue // Skip empty lines and comments
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid dictionary entry in %q: %s", filePath, line)
		}

		singlish := strings.TrimSpace(parts[0])
		goKeyword := strings.TrimSpace(parts[1])
		dict.mapping[singlish] = goKeyword

		// First entry wins as canonical for reverse lookup
		if _, exists := dict.reverseMapping[goKeyword]; !exists {
			dict.reverseMapping[goKeyword] = singlish
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading dictionary file %q: %w", filePath, err)
	}

	return dict, nil
}

// Lookup returns the Go keyword for a given Singlish keyword.
// It also returns a boolean indicating whether the keyword was found.
func (d *Dictionary) Lookup(singlishKeyword string) (string, bool) {
	goKeyword, found := d.mapping[singlishKeyword]
	return goKeyword, found
}

// ReverseLookup returns the canonical Singlish keyword for a given Go keyword.
func (d *Dictionary) ReverseLookup(goKeyword string) (string, bool) {
	singlish, found := d.reverseMapping[goKeyword]
	return singlish, found
}

// Keys returns all Singlish keywords in the dictionary.
func (d *Dictionary) Keys() []string {
	keys := make([]string, 0, len(d.mapping))
	for k := range d.mapping {
		keys = append(keys, k)
	}
	return keys
}
