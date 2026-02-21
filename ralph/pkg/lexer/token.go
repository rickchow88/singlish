package lexer

// TokenType describes the kind of token scanned from source.
type TokenType string

const (
	TokenIdentifier  TokenType = "identifier"
	TokenKeyword     TokenType = "keyword"
	TokenOperator    TokenType = "operator"
	TokenPunctuation TokenType = "punctuation"
	TokenString      TokenType = "string"
	TokenComment     TokenType = "comment"
	TokenNumber      TokenType = "number"
)

// Token represents a lexed token with its source location.
type Token struct {
	Type  TokenType
	Value string
	Line  int
	Col   int
}

// Diagnostic captures lexer errors with source location.
type Diagnostic struct {
	Message string
	Line    int
	Col     int
	Length  int
}
