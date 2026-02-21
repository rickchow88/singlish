package lexer

import "testing"

func TestLexIdentifiersAndKeywords(t *testing.T) {
	keywords := map[string]struct{}{"kampung": {}}
	input := "kampung boss"
	tokens, diagnostics := Lex(input, keywords)
	if len(diagnostics) != 0 {
		t.Fatalf("expected no diagnostics, got %v", diagnostics)
	}
	if len(tokens) != 2 {
		t.Fatalf("expected 2 tokens, got %d", len(tokens))
	}
	if tokens[0].Type != TokenKeyword || tokens[0].Value != "kampung" || tokens[0].Line != 1 || tokens[0].Col != 1 {
		t.Fatalf("unexpected first token: %#v", tokens[0])
	}
	if tokens[1].Type != TokenIdentifier || tokens[1].Value != "boss" || tokens[1].Line != 1 || tokens[1].Col != 9 {
		t.Fatalf("unexpected second token: %#v", tokens[1])
	}
}

func TestLexCommentsPreserveLineCounts(t *testing.T) {
	input := "boss // hi\n" +
		"kampung\n" +
		"/* multi\nline */\n" +
		"lagi"
	tokens, diagnostics := Lex(input, nil)
	if len(diagnostics) != 0 {
		t.Fatalf("expected no diagnostics, got %v", diagnostics)
	}
	if len(tokens) < 4 {
		t.Fatalf("expected at least 4 tokens, got %d", len(tokens))
	}
	if tokens[0].Value != "boss" || tokens[0].Line != 1 || tokens[0].Col != 1 {
		t.Fatalf("unexpected token[0]: %#v", tokens[0])
	}
	if tokens[1].Type != TokenComment || tokens[1].Line != 1 || tokens[1].Col != 6 {
		t.Fatalf("unexpected token[1]: %#v", tokens[1])
	}
	if tokens[2].Value != "kampung" || tokens[2].Line != 2 || tokens[2].Col != 1 {
		t.Fatalf("unexpected token[2]: %#v", tokens[2])
	}
	if tokens[3].Type != TokenComment || tokens[3].Line != 3 || tokens[3].Col != 1 {
		t.Fatalf("unexpected token[3]: %#v", tokens[3])
	}
	last := tokens[len(tokens)-1]
	if last.Value != "lagi" || last.Line != 5 || last.Col != 1 {
		t.Fatalf("unexpected last token: %#v", last)
	}
}

func TestLexStringLiterals(t *testing.T) {
	input := "gong(\"hi\")\n" +
		"gong(`raw\ntext`)"
	tokens, diagnostics := Lex(input, nil)
	if len(diagnostics) != 0 {
		t.Fatalf("expected no diagnostics, got %v", diagnostics)
	}
	var strings []Token
	for _, tok := range tokens {
		if tok.Type == TokenString {
			strings = append(strings, tok)
		}
	}
	if len(strings) != 2 {
		t.Fatalf("expected 2 string tokens, got %d", len(strings))
	}
	if strings[0].Value != "\"hi\"" || strings[0].Line != 1 || strings[0].Col != 6 {
		t.Fatalf("unexpected first string token: %#v", strings[0])
	}
	if strings[1].Value != "`raw\ntext`" || strings[1].Line != 2 || strings[1].Col != 6 {
		t.Fatalf("unexpected second string token: %#v", strings[1])
	}
}

func TestLexUnterminatedStringDiagnostic(t *testing.T) {
	input := "gong(\"hi\n"
	_, diagnostics := Lex(input, nil)
	if len(diagnostics) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diagnostics))
	}
	diag := diagnostics[0]
	if diag.Line != 1 || diag.Col != 6 {
		t.Fatalf("unexpected diagnostic location: %#v", diag)
	}
}
