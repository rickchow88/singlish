package lexer

import "unicode"

// Lex scans source input into tokens and diagnostics.
func Lex(input string, keywords map[string]struct{}) ([]Token, []Diagnostic) {
	l := newLexer(input, keywords)
	l.lex()
	return l.tokens, l.diagnostics
}

type lexer struct {
	src         []rune
	pos         int
	line        int
	col         int
	keywords    map[string]struct{}
	tokens      []Token
	diagnostics []Diagnostic
}

func newLexer(input string, keywords map[string]struct{}) *lexer {
	if keywords == nil {
		keywords = map[string]struct{}{}
	}
	return &lexer{
		src:      []rune(input),
		line:     1,
		col:      1,
		keywords: keywords,
	}
}

func (l *lexer) lex() {
	for !l.eof() {
		ch := l.peek()
		if isWhitespace(ch) {
			l.consumeWhitespace()
			continue
		}

		if ch == '/' {
			next := l.peekNext()
			if next == '/' {
				l.lexLineComment()
				continue
			}
			if next == '*' {
				if !l.lexBlockComment() {
					return
				}
				continue
			}
		}

		if ch == '"' || ch == '`' {
			if !l.lexString(ch) {
				return
			}
			continue
		}

		if isIdentifierStart(ch) {
			l.lexIdentifier()
			continue
		}

		if unicode.IsDigit(ch) {
			l.lexNumber()
			continue
		}

		if l.lexOperatorOrPunct() {
			continue
		}

		l.addDiagnostic("unexpected character", l.line, l.col)
		l.advance()
	}
}

func (l *lexer) eof() bool {
	return l.pos >= len(l.src)
}

func (l *lexer) peek() rune {
	if l.eof() {
		return 0
	}
	return l.src[l.pos]
}

func (l *lexer) peekNext() rune {
	if l.pos+1 >= len(l.src) {
		return 0
	}
	return l.src[l.pos+1]
}

func (l *lexer) advance() rune {
	if l.eof() {
		return 0
	}
	ch := l.src[l.pos]
	l.pos++
	if ch == '\n' {
		l.line++
		l.col = 1
		return ch
	}
	if ch == '\r' {
		if !l.eof() && l.src[l.pos] == '\n' {
			l.pos++
		}
		l.line++
		l.col = 1
		return ch
	}
	l.col++
	return ch
}

func (l *lexer) consumeWhitespace() {
	for !l.eof() && isWhitespace(l.peek()) {
		l.advance()
	}
}

func (l *lexer) lexIdentifier() {
	startLine, startCol := l.line, l.col
	startPos := l.pos
	l.advance()
	for !l.eof() && isIdentifierPart(l.peek()) {
		l.advance()
	}
	value := string(l.src[startPos:l.pos])
	tokType := TokenIdentifier
	if _, ok := l.keywords[value]; ok {
		tokType = TokenKeyword
	}
	l.tokens = append(l.tokens, Token{Type: tokType, Value: value, Line: startLine, Col: startCol})
}

func (l *lexer) lexNumber() {
	startLine, startCol := l.line, l.col
	startPos := l.pos
	l.advance()
	for !l.eof() && unicode.IsDigit(l.peek()) {
		l.advance()
	}
	if !l.eof() && l.peek() == '.' && l.pos+1 < len(l.src) && unicode.IsDigit(l.src[l.pos+1]) {
		l.advance()
		for !l.eof() && unicode.IsDigit(l.peek()) {
			l.advance()
		}
	}
	value := string(l.src[startPos:l.pos])
	l.tokens = append(l.tokens, Token{Type: TokenNumber, Value: value, Line: startLine, Col: startCol})
}

func (l *lexer) lexLineComment() {
	startLine, startCol := l.line, l.col
	startPos := l.pos
	l.advance()
	l.advance()
	for !l.eof() {
		ch := l.peek()
		if ch == '\n' || ch == '\r' {
			break
		}
		l.advance()
	}
	value := string(l.src[startPos:l.pos])
	l.tokens = append(l.tokens, Token{Type: TokenComment, Value: value, Line: startLine, Col: startCol})
}

func (l *lexer) lexBlockComment() bool {
	startLine, startCol := l.line, l.col
	startPos := l.pos
	l.advance()
	l.advance()
	for !l.eof() {
		if l.peek() == '*' && l.peekNext() == '/' {
			l.advance()
			l.advance()
			value := string(l.src[startPos:l.pos])
			l.tokens = append(l.tokens, Token{Type: TokenComment, Value: value, Line: startLine, Col: startCol})
			return true
		}
		l.advance()
	}
	l.addDiagnostic("unterminated block comment", startLine, startCol)
	return false
}

func (l *lexer) lexString(quote rune) bool {
	startLine, startCol := l.line, l.col
	startPos := l.pos
	l.advance()
	for !l.eof() {
		ch := l.peek()
		if quote == '`' {
			if ch == '`' {
				l.advance()
				value := string(l.src[startPos:l.pos])
				l.tokens = append(l.tokens, Token{Type: TokenString, Value: value, Line: startLine, Col: startCol})
				return true
			}
			l.advance()
			continue
		}

		if ch == '\\' {
			l.advance()
			if !l.eof() {
				l.advance()
			}
			continue
		}

		if ch == '"' {
			l.advance()
			value := string(l.src[startPos:l.pos])
			l.tokens = append(l.tokens, Token{Type: TokenString, Value: value, Line: startLine, Col: startCol})
			return true
		}

		if ch == '\n' || ch == '\r' {
			l.addDiagnostic("unterminated string literal", startLine, startCol)
			return false
		}

		l.advance()
	}

	l.addDiagnostic("unterminated string literal", startLine, startCol)
	return false
}

func (l *lexer) lexOperatorOrPunct() bool {
	startLine, startCol := l.line, l.col
	startPos := l.pos

	if match := l.matchOperator(); match != "" {
		l.tokens = append(l.tokens, Token{Type: TokenOperator, Value: match, Line: startLine, Col: startCol})
		return true
	}

	ch := l.peek()
	if isPunctuation(ch) {
		l.advance()
		value := string(l.src[startPos:l.pos])
		l.tokens = append(l.tokens, Token{Type: TokenPunctuation, Value: value, Line: startLine, Col: startCol})
		return true
	}

	return false
}

func (l *lexer) matchOperator() string {
	for _, op := range operators {
		runes := []rune(op)
		if l.pos+len(runes) > len(l.src) {
			continue
		}
		match := true
		for i, r := range runes {
			if l.src[l.pos+i] != r {
				match = false
				break
			}
		}
		if match {
			for range runes {
				l.advance()
			}
			return op
		}
	}
	return ""
}

func (l *lexer) addDiagnostic(message string, line, col int) {
	l.diagnostics = append(l.diagnostics, Diagnostic{Message: message, Line: line, Col: col})
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isIdentifierStart(ch rune) bool {
	return ch == '_' || unicode.IsLetter(ch)
}

func isIdentifierPart(ch rune) bool {
	return ch == '_' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func isPunctuation(ch rune) bool {
	switch ch {
	case '(', ')', '{', '}', '[', ']', ',', ';', ':', '.':
		return true
	default:
		return false
	}
}

var operators = []string{
	"...",
	"<<=",
	">>=",
	"&^=",
	"==",
	"!=",
	"<=",
	">=",
	"<-",
	"&&",
	"||",
	"++",
	"--",
	"+=",
	"-=",
	"*=",
	"/=",
	"%=",
	"&=",
	"|=",
	"^=",
	"<<",
	">>",
	"&^",
	":=",
	"=",
	"+",
	"-",
	"*",
	"/",
	"%",
	"<",
	">",
	"!",
	"&",
	"|",
	"^",
}
