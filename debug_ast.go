package main

import (
	"fmt"
	"os"

	"github.com/rickchow/singlish/pkg/dictionaries"
	"github.com/rickchow/singlish/pkg/lexer"
	"github.com/rickchow/singlish/pkg/parser"
)

func main() {
	input, _ := os.ReadFile("examples/53_type_assertion.singlish")
	dict := dictionaries.New()
	lex := lexer.New(string(input))
	lex.SetDictionary(dict)
	tokens := lex.Tokenize()
	p := parser.New(tokens, dict)
	p.ParseProgram()
	for _, e := range p.Errors() {
		fmt.Printf("Line %d:%d: %s\n", e.Line, e.Col, e.Message)
	}
}
