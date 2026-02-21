package dictionaries

// GetDefaultMappings returns the built-in Singlish to Go keyword mappings.
// This allows the transpiler to run without an external dictionary file.
func GetDefaultMappings() map[string]string {
	return map[string]string{
		"kampung": "package",
		"dapao":   "import",
		"action":  "func",
		"boss":    "main",

		"got":     "var",
		"confirm": "const",
		"auto":    "iota",
		"pattern": "type",

		"nasi":    "if",
		"den":     "else",
		"tikam":   "select",
		"see_how": "switch",
		"say":     "case",
		"tompang": "fallthrough",
		"anyhow":  "default",
		"flykite": "goto",

		"loop":  "for",
		"all":   "range",
		"cabut": "break",
		"go":    "continue",

		"balek": "return",
		"nanti": "defer",

		"chiong": "go",
		"lobang": "chan",
		"pass":   "<-",
		"catch":  "<-",

		"can":      "true",
		"cannot":   "false",
		"kosong":   "nil",
		"bolehtak": "bool",
		"nombor":   "int",
		"banyak":   "int64",
		"point":    "float64",
		"cheem":    "complex128",
		"tar":      "string",
		"barang":   "struct",
		"salah":    "error",
		"gabra":    "panic",
		"ki":       "*",
		"zhi":      "rune",
		"heng":     "recover",

		"kaki": "interface",
		"menu": "map",

		"buat":   "make",
		"upsize": "append",
		"buang":  "delete",
		"count":  "len",
		"kwear":  "close",
		"gong":   "fmt.Println",

		"somemore": "&&",
		"dun":      "!",
		"or":       "||",
	}
}

// NewDefaultDictionary creates a new Dictionary populated with the default Singlish mappings.
func NewDefaultDictionary() *Dictionary {
	defaults := GetDefaultMappings()
	dict := &Dictionary{
		mapping:        make(map[string]string),
		reverseMapping: make(map[string]string),
	}

	for singlish, goKeyword := range defaults {
		dict.mapping[singlish] = goKeyword

		// First entry wins as canonical for reverse lookup
		if _, exists := dict.reverseMapping[goKeyword]; !exists {
			dict.reverseMapping[goKeyword] = singlish
		}
	}
	return dict
}
