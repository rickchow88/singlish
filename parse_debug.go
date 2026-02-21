package main

import (
	"fmt"
	"os"

	"github.com/rickchow/singlish/pkg/codegen"
	"github.com/rickchow/singlish/pkg/dictionaries"
	"github.com/rickchow/singlish/pkg/lexer"
	"github.com/rickchow/singlish/pkg/parser"
)

func main() {
	input, _ := os.ReadFile("examples/87_sync_pool.singlish")
	dict := dictionaries.NewDefaultDictionary()

	keywords := make(map[string]struct{})
	for _, k := range dict.Keys() {
		keywords[k] = struct{}{}
	}
	keywords["ki"] = struct{}{}

	tokens, _ := lexer.Lex(string(input), keywords)
	p := parser.New(tokens, dict)
	program := p.ParseProgram()

	code, _ := codegen.Generate(program, dict)
	fmt.Println(code)
}
