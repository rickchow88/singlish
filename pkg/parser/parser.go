package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rickchow/singlish/pkg/ast"
	"github.com/rickchow/singlish/pkg/dictionaries"
	"github.com/rickchow/singlish/pkg/lexer"
)

const (
	_ int = iota
	LOWEST
	ASSIGN      // = or :=
	SEND        // <-
	LOGICAL_OR  // ||
	LOGICAL_AND // &&
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index] or struct.field
	POSTFIX     // i++
)

var precedence = map[string]int{
	"=":  ASSIGN,
	":=": ASSIGN,
	"+=": ASSIGN,
	"-=": ASSIGN,
	"*=": ASSIGN,
	"/=": ASSIGN,
	"%=": ASSIGN,
	"&=": ASSIGN,
	"|=": ASSIGN,
	"^=": ASSIGN,
	"<-": SEND,
	"==": EQUALS,
	"!=": EQUALS,
	"<":  LESSGREATER,
	">":  LESSGREATER,
	"<=": LESSGREATER,
	">=": LESSGREATER,
	"&&": LOGICAL_AND, // Should be logical AND
	"||": LOGICAL_OR,  // Should be logical OR
	"+":  SUM,
	"-":  SUM,
	"*":  PRODUCT,
	"/":  PRODUCT,
	"%":  PRODUCT,
	"(":  CALL,
	"[":  INDEX,
	"{":  CALL,
	".":  INDEX,
	"++": POSTFIX,
	"--": POSTFIX,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	tokens []lexer.Token
	dict   *dictionaries.Dictionary
	pos    int
	errors []lexer.Diagnostic

	curToken  lexer.Token
	peekToken lexer.Token

	prefixParseFns map[lexer.TokenType]prefixParseFn
	infixParseFns  map[lexer.TokenType]infixParseFn

	noCompositeLiteral bool
}

func New(tokens []lexer.Token, dict *dictionaries.Dictionary) *Parser {
	p := &Parser{
		tokens: tokens,
		dict:   dict,
		pos:    0,
		errors: []lexer.Diagnostic{},
	}

	p.prefixParseFns = make(map[lexer.TokenType]prefixParseFn)
	p.registerPrefix(lexer.TokenIdentifier, p.parseIdentifier)
	p.registerPrefix(lexer.TokenKeyword, p.parseIdentifier) // Allow keywords to be parsed as identifiers
	p.registerPrefix(lexer.TokenNumber, p.parseNumberLiteral)
	p.registerPrefix(lexer.TokenString, p.parseStringLiteral)
	p.registerPrefix(lexer.TokenOperator, p.parsePrefixExpression)     // for - and ! and <-
	p.registerPrefix(lexer.TokenPunctuation, p.parseGroupedExpression) // for (

	p.infixParseFns = make(map[lexer.TokenType]infixParseFn)
	p.registerInfix(lexer.TokenOperator, p.parseInfixExpression)
	p.registerInfix(lexer.TokenOperator, p.parseInfixExpression)
	p.registerInfix(lexer.TokenPunctuation, p.parseCallOrGroupExpression) // for ( or .
	p.registerInfix(lexer.TokenKeyword, p.parseInfixExpression)           // for keywords acting as operators (e.g. pass -> <-)
	p.registerInfix(lexer.TokenIdentifier, p.parseInfixExpression)        // for identifiers acting as operators

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	for {
		if p.pos < len(p.tokens) {
			p.peekToken = p.tokens[p.pos]
			p.pos++
		} else {
			p.peekToken = lexer.Token{Type: "EOF", Value: ""}
		}

		if p.peekToken.Type != lexer.TokenComment {
			break
		}
	}
}

func (p *Parser) curTokenIs(t lexer.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekPrecedence() int {
	val := p.peekToken.Value
	if p.dict != nil {
		if v, ok := p.dict.Lookup(val); ok {
			val = v
		}
	}
	if p, ok := precedence[val]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	val := p.curToken.Value
	if p.dict != nil {
		if v, ok := p.dict.Lookup(val); ok {
			val = v
		}
	}
	if p, ok := precedence[val]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) Errors() []lexer.Diagnostic {
	return p.errors
}

func (p *Parser) peekError(t lexer.TokenType, expectedValue string) {
	msg := fmt.Sprintf("expected next token to be %s (%s), got %s (%s) instead", t, expectedValue, p.peekToken.Type, p.peekToken.Value)
	p.errors = append(p.errors, lexer.Diagnostic{
		Message: msg,
		Line:    p.peekToken.Line,
		Col:     p.peekToken.Col,
		Length:  len(p.peekToken.Value),
	})
}

func (p *Parser) registerPrefix(tokenType lexer.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType lexer.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != "EOF" {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	tokenValue := p.curToken.Value
	canonical := tokenValue
	if p.dict != nil {
		if val, found := p.dict.Lookup(tokenValue); found {
			canonical = val
		}
	}

	switch canonical {
	case "package":
		s := p.parsePackageStatement()
		if s == nil {
			return nil
		}
		return s
	case "import":
		s := p.parseImportStatement()
		if s == nil {
			return nil
		}
		return s
	case "var", "const", "let":
		s := p.parseLetStatement()
		if s == nil {
			return nil
		}
		return s
	case "return":
		s := p.parseReturnStatement()
		if s == nil {
			return nil
		}
		return s
	case "func":
		s := p.parseFunctionStatement()
		if s == nil {
			return nil
		}
		return s
	case "type":
		s := p.parseTypeStatement()
		if s == nil {
			return nil
		}
		return s
	case "if":
		s := p.parseIfStatement()
		if s == nil {
			return nil
		}
		return s
	case "for":
		s := p.parseForStatement()
		if s == nil {
			return nil
		}
		return s
	case "go":
		s := p.parseGoStatement()
		if s == nil {
			return nil
		}
		return s
	case "defer":
		s := p.parseDeferStatement()
		if s == nil {
			return nil
		}
		return s
	case "switch":
		s := p.parseSwitchStatement()
		if s == nil {
			return nil
		}
		return s
	case "select":
		s := p.parseSelectStatement()
		if s == nil {
			return nil
		}
		return s
	default:
		s := p.parseExpressionStatement()
		if s == nil {
			return nil
		}
		return s
	}
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curToken}

	p.nextToken() // skip 'if'

	p.noCompositeLiteral = true
	stmt.Condition = p.parseExpression(LOWEST)
	p.noCompositeLiteral = false

	if !p.expectPeek(lexer.TokenPunctuation, "{") {
		return nil
	}

	stmt.Consequence = p.parseBlockStatement()

	if p.peekCanonical("else") {
		p.nextToken() // consume 'else'

		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "{" {
			p.nextToken() // consume '{'
			stmt.AlternativeStmt = p.parseBlockStatement()
		} else if p.peekCanonical("if") {
			// else if ...
			p.nextToken() // consume 'if'
			stmt.AlternativeStmt = p.parseIfStatement()
		}
	}

	return stmt
}

func (p *Parser) parseGoStatement() *ast.GoStatement {
	stmt := &ast.GoStatement{Token: p.curToken}
	p.nextToken() // skip 'go'

	exp := p.parseExpression(LOWEST)
	if call, ok := exp.(*ast.CallExpression); ok {
		stmt.Call = call
		return stmt
	}

	// Error if not a call
	msg := fmt.Sprintf("expected function call after go, got %T", exp)
	p.errors = append(p.errors, lexer.Diagnostic{
		Message: msg,
		Line:    stmt.Token.Line,
		Col:     stmt.Token.Col,
		Length:  len(stmt.Token.Value),
	})
	return nil
}

func (p *Parser) parseDeferStatement() *ast.DeferStatement {
	stmt := &ast.DeferStatement{Token: p.curToken}
	p.nextToken() // skip 'defer'

	exp := p.parseExpression(LOWEST)
	if call, ok := exp.(*ast.CallExpression); ok {
		stmt.Call = call
		return stmt
	}

	msg := fmt.Sprintf("expected function call after defer, got %T", exp)
	p.errors = append(p.errors, lexer.Diagnostic{
		Message: msg,
		Line:    stmt.Token.Line,
		Col:     stmt.Token.Col,
		Length:  len(stmt.Token.Value),
	})
	return nil
}

func (p *Parser) peekCanonical(kw string) bool {
	val := p.peekToken.Value
	if p.dict != nil {
		if canonical, ok := p.dict.Lookup(val); ok {
			val = canonical
		}
	}
	return val == kw
}

func (p *Parser) curCanonical(kw string) bool {
	val := p.curToken.Value
	if p.dict != nil {
		if canonical, ok := p.dict.Lookup(val); ok {
			val = canonical
		}
	}
	return val == kw
}

func (p *Parser) parseForStatement() *ast.ForStatement {
	stmt := &ast.ForStatement{Token: p.curToken}
	p.nextToken() // skip 'for'

	// Infinite loop: for { ... }
	if p.curTokenIs(lexer.TokenPunctuation) && p.curToken.Value == "{" {
		stmt.Body = p.parseBlockStatement()
		return stmt
	}

	// Potential Range loop: loop [i, v] = all orders { ... }
	if p.curTokenIs(lexer.TokenIdentifier) || p.curTokenIs(lexer.TokenKeyword) {
		backupPos := p.pos
		backupCur := p.curToken
		backupPeek := p.peekToken

		idents := p.parseIdentifierList()
		if idents != nil && (p.curToken.Value == "=" || p.curToken.Value == ":=") && p.peekCanonical("range") {
			stmt.IsRange = true
			if len(idents) > 0 {
				stmt.Key = idents[0]
			}
			if len(idents) > 1 {
				stmt.Value = idents[1]
			}
			p.nextToken() // consume =
			p.nextToken() // consume all
			p.noCompositeLiteral = true
			stmt.Iterable = p.parseExpression(LOWEST)
			p.noCompositeLiteral = false

			if !p.expectPeek(lexer.TokenPunctuation, "{") {
				return nil
			}
			stmt.Body = p.parseBlockStatement()
			return stmt
		}

		p.pos = backupPos
		p.curToken = backupCur
		p.peekToken = backupPeek
	}

	// 1. "for init; cond; post {"
	p.noCompositeLiteral = true
	first := p.parseStatement()
	p.noCompositeLiteral = false

	if p.curTokenIs(lexer.TokenPunctuation) && p.curToken.Value == ";" {
		stmt.Init = first

		if !p.peekTokenIs(lexer.TokenPunctuation) || p.peekToken.Value != ";" {
			p.nextToken() // advance past the semicolon
			p.noCompositeLiteral = true
			stmt.Condition = p.parseExpression(LOWEST)
			p.noCompositeLiteral = false
		}

		if !p.expectPeek(lexer.TokenPunctuation, ";") {
			return nil
		}

		if !p.peekTokenIs(lexer.TokenPunctuation) || p.peekToken.Value != "{" {
			p.nextToken() // advance past the semicolon
			stmt.Post = p.parseStatement()
		}

		if !p.expectPeek(lexer.TokenPunctuation, "{") {
			return nil
		}
		stmt.Body = p.parseBlockStatement()
		return stmt
	}

	// 2. "for condition {"
	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "{" {
		if es, ok := first.(*ast.ExpressionStatement); ok {
			stmt.Condition = es.Expression
		}
		p.nextToken() // skip to {
		stmt.Body = p.parseBlockStatement()
		return stmt
	}

	return nil
}

func (p *Parser) parseIdentifierList() []*ast.Identifier {
	list := []*ast.Identifier{}

	if !p.curTokenIs(lexer.TokenIdentifier) && !p.curTokenIs(lexer.TokenKeyword) {
		return nil
	}

	list = append(list, &ast.Identifier{Token: p.curToken, Value: p.curToken.Value})

	for p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "," {
		p.nextToken() // consume ,
		if !p.expectPeekType(lexer.TokenIdentifier) && !p.peekTokenIs(lexer.TokenKeyword) {
			return nil
		}
		list = append(list, &ast.Identifier{Token: p.curToken, Value: p.curToken.Value})
	}

	p.nextToken()
	return list
}

func (p *Parser) parsePostfixExpression(left ast.Expression) ast.Expression {
	if p.curToken.Value == "..." {
		return &ast.Identifier{Token: p.curToken, Value: left.String() + "..."}
	}
	return &ast.IncDecStatement{
		Token:    p.curToken,
		Left:     left,
		Operator: p.curToken.Value,
	}
}

func (p *Parser) parsePackageStatement() *ast.PackageStatement {
	stmt := &ast.PackageStatement{Token: p.curToken}

	if !p.expectPeekType(lexer.TokenIdentifier) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}
	return stmt
}

func (p *Parser) parseImportStatement() *ast.ImportStatement {
	stmt := &ast.ImportStatement{Token: p.curToken}

	if !p.expectPeekType(lexer.TokenString) {
		return nil
	}

	stmt.Path = &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Value}
	return stmt
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	idents := p.parseIdentifierListForLet()
	if idents == nil {
		return nil
	}
	stmt.Names = idents

	// Check for type annotation
	// Type can be: identifier, keyword (like nombor), or [ for array/slice
	if p.peekTokenIs(lexer.TokenIdentifier) || p.peekTokenIs(lexer.TokenKeyword) ||
		(p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "[") ||
		(p.peekTokenIs(lexer.TokenOperator) && p.peekToken.Value == "*") {
		stmt.Type = p.parseType()
	}

	if p.peekTokenIs(lexer.TokenOperator) && (p.peekToken.Value == "=" || p.peekToken.Value == ":=") {
		p.nextToken()
		p.nextToken()
		stmt.Value = p.parseExpression(LOWEST)
	}

	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ";" {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseIdentifierListForLet() []*ast.Identifier {
	list := []*ast.Identifier{}

	if !p.expectPeekType(lexer.TokenIdentifier) && !p.peekTokenIs(lexer.TokenKeyword) {
		return nil
	}

	list = append(list, &ast.Identifier{Token: p.curToken, Value: p.curToken.Value})

	for p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "," {
		p.nextToken() // ,
		if !p.expectPeekType(lexer.TokenIdentifier) && !p.peekTokenIs(lexer.TokenKeyword) {
			return nil
		}
		list = append(list, &ast.Identifier{Token: p.curToken, Value: p.curToken.Value})
	}

	return list
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// Handle explicit return without value (if followed by ; or EOF)
	if (p.curTokenIs(lexer.TokenPunctuation) && p.curToken.Value == ";") || p.curTokenIs("EOF") {
		return stmt
	}

	// Also stop if current token is a keyword that starts a new statement/case
	if p.curCanonical("say") || p.curCanonical("case") || p.curCanonical("anyhow") || p.curCanonical("default") {
		return stmt
	}

	// Parse comma-separated expressions
	for {

		stmt.ReturnValues = append(stmt.ReturnValues, p.parseExpression(LOWEST))
		if !p.peekTokenIs(lexer.TokenPunctuation) || p.peekToken.Value != "," {
			break
		}
		p.nextToken() // ,
		p.nextToken() // next expression
	}

	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ";" {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {

	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ";" {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseTypeStatement() *ast.TypeStatement {
	stmt := &ast.TypeStatement{Token: p.curToken}

	if !p.expectPeekType(lexer.TokenIdentifier) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}

	if p.peekTokenIs(lexer.TokenOperator) && p.peekToken.Value == "=" {
		p.nextToken() // consume '='
		stmt.IsAlias = true
	}

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs("EOF") && precedence < p.peekPrecedence() {
		// Special case: Do not parse '{' (CompositeLiteral) if flag is set (e.g. in if/for condition)
		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "{" {
			if p.noCompositeLiteral {
				return leftExp
			}

			if _, ok := leftExp.(*ast.IncDecStatement); ok {
				return leftExp
			}
			if infix, ok := leftExp.(*ast.InfixExpression); ok {
				if infix.Operator != "." {
					return leftExp
				}
			}
			if prefix, ok := leftExp.(*ast.PrefixExpression); ok {
				// Allow &T{} and *T{}
				if prefix.Operator != "&" && prefix.Operator != "*" {
					return leftExp
				}
			}
			if _, ok := leftExp.(*ast.IndexExpression); ok {
				return leftExp
			}
			if _, ok := leftExp.(*ast.CallExpression); ok {
				return leftExp
			}
			if _, ok := leftExp.(*ast.IntegerLiteral); ok {
				return leftExp
			}
			if _, ok := leftExp.(*ast.FloatLiteral); ok {
				return leftExp
			}
			if _, ok := leftExp.(*ast.StringLiteral); ok {
				return leftExp
			}
			// Block literal values (true, false, nil) from being Types for CompositeLiterals
			if ident, ok := leftExp.(*ast.Identifier); ok {
				if ident.Value == "true" || ident.Value == "false" || ident.Value == "nil" {
					return leftExp
				}
			}
		}

		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	canonical := p.curToken.Value
	if p.dict != nil {
		if val, found := p.dict.Lookup(p.curToken.Value); found {
			canonical = val
		}
	}

	if canonical == "struct" {
		return p.parseStructLiteral()
	}
	if canonical == "interface" {
		return p.parseInterfaceLiteral()
	}

	// Handle 'func' (action) as closure
	if canonical == "func" {
		return p.parseFunctionLiteral()
	}

	// Handle pointer dereference (ki -> *)
	if canonical == "*" {
		expression := &ast.PrefixExpression{
			Token:    p.curToken,
			Operator: canonical,
		}
		p.nextToken()
		expression.Right = p.parseExpression(PREFIX)
		return expression
	}

	// Handle channel receive (catch -> <-)
	if canonical == "<-" {
		expression := &ast.PrefixExpression{
			Token:    p.curToken,
			Operator: canonical,
		}
		p.nextToken()
		expression.Right = p.parseExpression(PREFIX)
		return expression
	}

	// Handle 'chan' (lobang) type used as expression (e.g. in make)
	if canonical == "chan" {
		val := canonical
		if p.peekTokenIs(lexer.TokenOperator) && p.peekToken.Value == "<" {
			p.nextToken() // <
			elem := p.parseType()
			if p.peekTokenIs(lexer.TokenOperator) && p.peekToken.Value == ">" {
				p.nextToken() // >
			}
			val = "chan " + elem.Value
		} else if p.peekTokenIs(lexer.TokenIdentifier) || p.peekTokenIs(lexer.TokenKeyword) {
			// Standard Go syntax: chan int
			// parseType expects to consume from peek.
			elem := p.parseType()
			val = "chan " + elem.Value
		}

		return &ast.Identifier{Token: p.curToken, Value: val}
	}

	// Handle 'map' (menu)
	if canonical == "map" {
		val := canonical
		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "[" {
			p.nextToken() // [
			key := p.parseType()
			if p.expectPeek(lexer.TokenPunctuation, "]") {
				// ]
			}
			elem := p.parseType()
			val = "map[" + key.Value + "]" + elem.Value
		}
		return &ast.Identifier{Token: p.curToken, Value: val}
	}

	return &ast.Identifier{Token: p.curToken, Value: canonical}
}

func (p *Parser) parseNumberLiteral() ast.Expression {
	val := p.curToken.Value
	if strings.Contains(val, ".") {
		return p.parseFloatLiteral()
	}
	return p.parseIntegerLiteral()
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Value, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Value)
		p.errors = append(p.errors, lexer.Diagnostic{
			Message: msg,
			Line:    p.curToken.Line,
			Col:     p.curToken.Col,
			Length:  len(p.curToken.Value),
		})
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{Token: p.curToken}
	value, err := strconv.ParseFloat(p.curToken.Value, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.curToken.Value)
		p.errors = append(p.errors, lexer.Diagnostic{
			Message: msg,
			Line:    p.curToken.Line,
			Col:     p.curToken.Col,
			Length:  len(p.curToken.Value),
		})
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Value}
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	if p.curToken.Value != "-" && p.curToken.Value != "!" && p.curToken.Value != "&" && p.curToken.Value != "*" && p.curToken.Value != "<-" {
		return nil
	}

	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Value,
	}

	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	if p.curToken.Value == "++" || p.curToken.Value == "--" {
		return p.parsePostfixExpression(left)
	}

	operator := p.curToken.Value
	if p.dict != nil {
		if val, found := p.dict.Lookup(operator); found {
			operator = val
		}
	}

	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: operator,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	if p.curToken.Value == "[" {
		// Handle slice/array type literal: []int or [5]int
		p.nextToken() // consume [

		// Empty slice []T
		if p.curTokenIs(lexer.TokenPunctuation) && p.curToken.Value == "]" {
			elemType := p.parseType()
			return &ast.Identifier{Token: p.curToken, Value: "[]" + elemType.Value}
		}

		// Array [N]T - simplistic skipping for now
		// parsing expression inside [] not fully supported for array size in types yet
		// Just consume until ]
		for !p.curTokenIs(lexer.TokenPunctuation) || p.curToken.Value != "]" {
			p.nextToken()
		}
		p.nextToken() // consume ]
		elemType := p.parseType()
		// Return generic array for now or construct value
		return &ast.Identifier{Token: p.curToken, Value: "[]" + elemType.Value}
	}

	if p.curToken.Value == "{" {
		return p.parseCompositeLiteral(nil)
	}

	if p.curToken.Value != "(" {
		return nil
	}

	p.nextToken()
	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.TokenPunctuation, ")") {
		return nil
	}

	return exp
}

func (p *Parser) parseCallOrGroupExpression(left ast.Expression) ast.Expression {
	if p.curToken.Value == "(" {
		return p.parseCallExpression(left)
	}
	if p.curToken.Value == "." {
		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "(" {
			exp := &ast.TypeAssertionExpression{
				Token: p.curToken,
				Left:  left,
			}
			p.nextToken() // move to '('
			exp.Type = p.parseTypeExpression()
			if !p.expectPeek(lexer.TokenPunctuation, ")") {
				return nil
			}
			return exp
		}
		return p.parseInfixExpression(left)
	}
	if p.curToken.Value == "[" {
		return p.parseIndexExpression(left)
	}
	if p.curToken.Value == "{" {
		return p.parseCompositeLiteral(left)
	}
	return nil
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	tok := p.curToken

	// Check if it's an empty slice a[:] or a[:x]
	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ":" {
		p.nextToken() // move to ':'
		se := &ast.SliceExpression{
			Token: tok,
			Left:  left,
		}
		// Is high bound empty? a[:]
		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "]" {
			if !p.expectPeek(lexer.TokenPunctuation, "]") {
				return nil
			}
			return se
		}
		// Otherwise, parse high bound
		p.nextToken() // move to start of high
		se.High = p.parseExpression(LOWEST)
		if !p.expectPeek(lexer.TokenPunctuation, "]") {
			return nil
		}
		return se
	}

	// We have a low bound or an exact index
	p.nextToken() // move to start of expression
	indexOrLow := p.parseExpression(LOWEST)

	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ":" {
		// It's a slice expression
		se := &ast.SliceExpression{
			Token: tok,
			Left:  left,
			Low:   indexOrLow,
		}
		p.nextToken() // move to ':'
		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "]" {
			// a[x:]
			if !p.expectPeek(lexer.TokenPunctuation, "]") {
				return nil
			}
			return se
		}
		// a[x:y]
		p.nextToken() // move to start of high
		se.High = p.parseExpression(LOWEST)
		if !p.expectPeek(lexer.TokenPunctuation, "]") {
			return nil
		}
		return se
	}

	// It's just an index expression
	if !p.expectPeek(lexer.TokenPunctuation, "]") {
		return nil
	}

	return &ast.IndexExpression{
		Token: tok,
		Left:  left,
		Index: indexOrLow,
	}
}

func (p *Parser) parseCompositeLiteral(left ast.Expression) ast.Expression {
	lit := &ast.CompositeLiteral{
		Token: p.curToken,
		Type:  left,
	}
	lit.Elements = p.parseCompositeLiteralElements()
	return lit
}

func (p *Parser) parseCompositeLiteralElements() []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "}" {
		p.nextToken()
		return list
	}

	p.nextToken()

	for {
		exp := p.parseExpression(LOWEST)

		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ":" {
			p.nextToken() // consume :
			p.nextToken() // move to value
			value := p.parseExpression(LOWEST)
			exp = &ast.KeyValueExpression{
				Token: lexer.Token{Type: lexer.TokenPunctuation, Value: ":", Line: p.curToken.Line, Col: p.curToken.Col},
				Key:   exp,
				Value: value,
			}
		}

		list = append(list, exp)

		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "}" {
			break
		}

		if !p.expectPeek(lexer.TokenPunctuation, ",") {
			return nil
		}

		// Allow trailing comma? if peek is }
		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "}" {
			break
		}

		p.nextToken()
	}

	if !p.expectPeek(lexer.TokenPunctuation, "}") {
		return nil
	}

	return list
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ")" {
		p.nextToken()
		return args
	}

	p.nextToken()
	arg := p.parseExpression(LOWEST)
	// Check for variadic expansion (slice...)
	if p.peekTokenIs(lexer.TokenOperator) && p.peekToken.Value == "..." {
		p.nextToken() // consume ...
		arg = &ast.Identifier{Token: p.curToken, Value: arg.String() + "..."}
	}
	args = append(args, arg)

	for p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "," {
		p.nextToken()
		p.nextToken()
		arg := p.parseExpression(LOWEST)
		// Check for variadic expansion
		if p.peekTokenIs(lexer.TokenOperator) && p.peekToken.Value == "..." {
			p.nextToken() // consume ...
			arg = &ast.Identifier{Token: p.curToken, Value: arg.String() + "..."}
		}
		args = append(args, arg)
	}

	if !p.expectPeek(lexer.TokenPunctuation, ")") {
		return nil
	}

	return args
}

func (p *Parser) expectPeek(t lexer.TokenType, val string) bool {
	if p.peekToken.Type == t && p.peekToken.Value == val {
		p.nextToken()
		return true
	}
	p.peekError(t, val)
	return false
}

func (p *Parser) expectPeekType(t lexer.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}
	p.peekError(t, "")
	return false
}

func (p *Parser) parseType() *ast.Identifier {
	// 1. Pointers (* / ki)
	if p.peekTokenIs(lexer.TokenKeyword) && p.peekToken.Value == "ki" {
		p.nextToken()
		sub := p.parseType()
		return &ast.Identifier{Token: p.curToken, Value: "*" + sub.Value}
	}
	if p.peekTokenIs(lexer.TokenOperator) && p.peekToken.Value == "*" {
		p.nextToken()
		sub := p.parseType()
		return &ast.Identifier{Token: p.curToken, Value: "*" + sub.Value}
	}

	// 2. Arrays and Slices - Check for [
	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "[" {
		p.nextToken() // [
		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "]" {
			p.nextToken() // ]
			// It's a slice
			sub := p.parseType()
			return &ast.Identifier{Token: p.curToken, Value: "[]" + sub.Value}
		}
		// Array: [n]T
		if p.peekTokenIs(lexer.TokenNumber) {
			p.nextToken() // size
			size := p.curToken.Value
			if !p.expectPeek(lexer.TokenPunctuation, "]") {
				return nil
			}
			sub := p.parseType()
			return &ast.Identifier{Token: p.curToken, Value: "[" + size + "]" + sub.Value}
		}
	}

	p.nextToken()
	val := p.curToken.Value
	canonical := val
	if p.dict != nil {
		if v, ok := p.dict.Lookup(val); ok {
			canonical = v
		}
	}

	// Handle pkg.Type
	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "." {
		p.nextToken() // .
		p.nextToken() // Type name
		canonical = canonical + "." + p.curToken.Value
	}

	// 3. Map (menu / map)
	if canonical == "map" {
		// Expect [ Key ] Value
		if !p.expectPeek(lexer.TokenPunctuation, "[") {
			return &ast.Identifier{Token: p.curToken, Value: canonical}
		}
		keyType := p.parseType()
		if !p.expectPeek(lexer.TokenPunctuation, "]") {
			return &ast.Identifier{Token: p.curToken, Value: canonical} // error?
		}
		valType := p.parseType()
		return &ast.Identifier{Token: p.curToken, Value: "map[" + keyType.Value + "]" + valType.Value}
	}

	// 4. Channel (lobang / chan)
	if canonical == "chan" {
		// Handle lobang<tar> (Singlish) or chan string (Go)
		if p.peekTokenIs(lexer.TokenOperator) && p.peekToken.Value == "<" {
			p.nextToken() // <
			elem := p.parseType()
			if p.peekTokenIs(lexer.TokenOperator) && p.peekToken.Value == ">" {
				p.nextToken() // >
			}
			return &ast.Identifier{Token: p.curToken, Value: "chan " + elem.Value}
		}

		// Standard chan Type
		elem := p.parseType()
		return &ast.Identifier{Token: p.curToken, Value: "chan " + elem.Value}
	}

	// Handle interface{} / kaki{}
	if canonical == "interface" {
		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "{" {
			p.nextToken() // {
			if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "}" {
				p.nextToken() // }
				return &ast.Identifier{Token: p.curToken, Value: "interface{}"}
			}
		}
	}

	return &ast.Identifier{Token: p.curToken, Value: canonical}
}

func (p *Parser) noPrefixParseFnError(t lexer.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, lexer.Diagnostic{
		Message: msg,
		Line:    p.curToken.Line,
		Col:     p.curToken.Col,
		Length:  len(p.curToken.Value),
	})
}

func (p *Parser) parseFunctionStatement() *ast.FunctionStatement {
	stmt := &ast.FunctionStatement{Token: p.curToken}

	// Check for Receiver if next token is '('
	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "(" {
		p.nextToken() // consume '('

		p.nextToken() // move to receiver name
		nameTok := p.curToken

		// Parse Receiver Type (handles ki/*)
		recvType := p.parseType()

		stmt.Receiver = &ast.FieldDefinition{
			Name: &ast.Identifier{Token: nameTok, Value: nameTok.Value},
			Type: recvType,
		}

		if !p.expectPeek(lexer.TokenPunctuation, ")") {
			return nil
		}
	}

	// Allow identifier OR keyword (for 'boss' -> 'main')
	if !p.peekTokenIs(lexer.TokenIdentifier) && !p.peekTokenIs(lexer.TokenKeyword) {
		p.peekError(lexer.TokenIdentifier, "identifier")
		return nil
	}
	p.nextToken()

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}

	if !p.expectPeek(lexer.TokenPunctuation, "(") {
		return nil
	}

	stmt.Parameters = p.parseFunctionParameters()

	// Check for Return Type
	if !p.peekTokenIs(lexer.TokenPunctuation) || p.peekToken.Value != "{" {
		// Might be a single type or a grouped list of types
		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "(" {
			// (tar, salah) or (a int, b int)
			p.nextToken() // (
			types := []ast.Expression{}
			for {
				t := p.parseTypeExpression()
				types = append(types, t)
				if !p.peekTokenIs(lexer.TokenPunctuation) || p.peekToken.Value != "," {
					break
				}
				p.nextToken() // ,
			}
			if !p.expectPeek(lexer.TokenPunctuation, ")") {
				return nil
			}
			// For now, represent grouped return types as a special structure or just a string representation
			// Simplified: just store as a list-like expression if needed, but AST currently has Expression field.
			// Wrap in a GroupedExpression?
			// Let's use a simple Identifier for the whole lot for now, or just the first type if multiple not handled.
			// Actually, let's just use the String() representation.
			vals := []string{}
			for _, t := range types {
				vals = append(vals, t.String())
			}
			stmt.ReturnType = &ast.Identifier{Token: p.curToken, Value: "(" + strings.Join(vals, ", ") + ")"}
		} else {
			stmt.ReturnType = p.parseTypeExpression()
		}
	}

	if !p.expectPeek(lexer.TokenPunctuation, "{") {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseTypeExpression() ast.Expression {
	// Variadic: check if next token is ...
	if p.peekTokenIs(lexer.TokenOperator) && p.peekToken.Value == "..." {
		p.nextToken() // move to ...
		sub := p.parseType()
		return &ast.Identifier{Token: p.curToken, Value: "..." + sub.Value}
	}
	return p.parseType()
}

func (p *Parser) parseSwitchStatement() *ast.SwitchStatement {
	stmt := &ast.SwitchStatement{Token: p.curToken}
	p.nextToken() // skip switch keyword

	// Parse optional switch expression (if not immediately followed by {)
	if !p.curTokenIs(lexer.TokenPunctuation) || p.curToken.Value != "{" {
		p.noCompositeLiteral = true
		stmt.Expression = p.parseExpression(LOWEST)
		p.noCompositeLiteral = false
	}

	if !p.expectPeek(lexer.TokenPunctuation, "{") {
		return nil
	}

	// Parse case statements
	for p.peekCanonical("say") || p.peekCanonical("case") || p.peekCanonical("anyhow") || p.peekCanonical("default") {
		stmt.Cases = append(stmt.Cases, p.parseCaseStatement())
	}

	if !p.expectPeek(lexer.TokenPunctuation, "}") {
		return nil
	}

	return stmt
}

func (p *Parser) parseCaseStatement() *ast.CaseStatement {
	p.nextToken() // consume case/default keyword
	stmt := &ast.CaseStatement{Token: p.curToken}

	canonical := p.curToken.Value
	if p.dict != nil {
		if v, ok := p.dict.Lookup(canonical); ok {
			canonical = v
		}
	}

	if canonical == "default" {
		stmt.Default = true
	} else {
		// Advance to the first expression
		p.nextToken()
		// Parse list of expressions
		for {
			stmt.Expressions = append(stmt.Expressions, p.parseExpression(LOWEST))
			if !p.peekTokenIs(lexer.TokenPunctuation) || p.peekToken.Value != "," {
				break
			}
			p.nextToken() // consume ,
			p.nextToken() // move to next expression
		}
	}

	if !p.expectPeek(lexer.TokenPunctuation, ":") {
		// Error already reported
		return stmt
	}

	// Parse body until next case or }
	stmt.Body = &ast.BlockStatement{Token: p.curToken}
	for !p.peekTokenIs(lexer.TokenPunctuation) || p.peekToken.Value != "}" {
		if p.peekCanonical("case") || p.peekCanonical("say") || p.peekCanonical("default") || p.peekCanonical("anyhow") {
			break
		}
		if p.peekTokenIs("EOF") {
			break
		}

		p.nextToken()
		s := p.parseStatement()
		if s != nil {
			stmt.Body.Statements = append(stmt.Body.Statements, s)
		}
	}

	return stmt
}

func (p *Parser) parseSelectStatement() *ast.SelectStatement {
	stmt := &ast.SelectStatement{Token: p.curToken}

	// select has no expression, just { cases }
	if !p.expectPeek(lexer.TokenPunctuation, "{") {
		return nil
	}

	// Parse select cases
	for p.peekCanonical("say") || p.peekCanonical("case") || p.peekCanonical("anyhow") || p.peekCanonical("default") {
		stmt.Cases = append(stmt.Cases, p.parseSelectCase())
	}

	if !p.expectPeek(lexer.TokenPunctuation, "}") {
		return nil
	}

	return stmt
}

func (p *Parser) parseSelectCase() *ast.SelectCase {
	p.nextToken() // consume case/default keyword
	stmt := &ast.SelectCase{Token: p.curToken}

	canonical := p.curToken.Value
	if p.dict != nil {
		if v, ok := p.dict.Lookup(canonical); ok {
			canonical = v
		}
	}

	if canonical == "default" {
		stmt.Default = true
	} else {
		// Parse communication clause: can be:
		// - <-c (receive and discard)
		// - x := <-c (receive with assignment)
		// - x = <-c (receive with assignment)
		// - c <- v (send)
		p.nextToken() // move to first token of comm clause

		// Parse as expression statement which handles both sends and receives
		comm := p.parseExpressionStatement()
		stmt.Comm = comm
	}

	if !p.expectPeek(lexer.TokenPunctuation, ":") {
		return stmt
	}

	// Parse body until next case or }
	stmt.Body = &ast.BlockStatement{Token: p.curToken}
	for !p.peekTokenIs(lexer.TokenPunctuation) || p.peekToken.Value != "}" {
		if p.peekCanonical("case") || p.peekCanonical("say") || p.peekCanonical("default") || p.peekCanonical("anyhow") {
			break
		}
		if p.peekTokenIs("EOF") {
			break
		}

		p.nextToken()
		s := p.parseStatement()
		if s != nil {
			stmt.Body.Statements = append(stmt.Body.Statements, s)
		}
	}

	return stmt
}

func (p *Parser) parseFunctionParameters() []*ast.FieldDefinition {

	identifiers := []*ast.FieldDefinition{}

	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ")" {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	for {
		nameTok := p.curToken

		ident := &ast.FieldDefinition{
			Name: &ast.Identifier{Token: nameTok, Value: nameTok.Value},
		}

		if p.peekTokenIs(lexer.TokenPunctuation) && (p.peekToken.Value == "," || p.peekToken.Value == ")") {
			ident.Type = nil
		} else {
			ident.Type = p.parseTypeExpression()
		}

		identifiers = append(identifiers, ident)

		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "," {
			p.nextToken() // move to ','
			p.nextToken() // move to next identifier
		} else {
			break
		}
	}

	if !p.expectPeek(lexer.TokenPunctuation, ")") {
		return nil
	}

	// Backtrack to apply types to grouped parameters
	var lastType ast.Expression
	for i := len(identifiers) - 1; i >= 0; i-- {
		if identifiers[i].Type != nil {
			lastType = identifiers[i].Type
		} else if lastType != nil {
			identifiers[i].Type = lastType
		} else {
			identifiers[i].Type = identifiers[i].Name
			identifiers[i].Name = nil
		}
	}

	return identifiers
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(lexer.TokenPunctuation) || p.curToken.Value != "}" {
		if p.curToken.Type == "EOF" {
			break
		}
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseStructLiteral() ast.Expression {
	lit := &ast.StructLiteral{Token: p.curToken}

	if !p.expectPeek(lexer.TokenPunctuation, "{") {
		return nil
	}

	lit.Fields = p.parseStructFields()

	return lit
}

func (p *Parser) parseStructFields() []*ast.FieldDefinition {
	fields := []*ast.FieldDefinition{}

	p.nextToken()

	for !p.curTokenIs(lexer.TokenPunctuation) || p.curToken.Value != "}" {
		if p.curToken.Type == "EOF" {
			break
		}

		field := &ast.FieldDefinition{}

		// Check canonical for 'var' (got) and skip it
		canonical := p.curToken.Value
		if p.dict != nil {
			if val, found := p.dict.Lookup(canonical); found {
				canonical = val
			}
		}

		if canonical == "var" || canonical == "got" {
			p.nextToken()
		}

		// Expect Identifier (Name)
		if !p.curTokenIs(lexer.TokenIdentifier) && !p.curTokenIs(lexer.TokenKeyword) {
			// Should validation error here?
			// For now, skip until }
			p.nextToken()
			continue
		}

		field.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}
		p.nextToken()

		// Parse Type
		field.Type = p.parseExpression(LOWEST)

		// Parse Tag (StringLiteral) if present
		if p.peekTokenIs(lexer.TokenString) {
			p.nextToken()
			field.Tag = &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Value}
		}

		fields = append(fields, field)

		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ";" {
			p.nextToken()
		}

		p.nextToken()
	}

	return fields
}

func (p *Parser) parseInterfaceLiteral() ast.Expression {
	lit := &ast.InterfaceLiteral{Token: p.curToken}

	if !p.expectPeek(lexer.TokenPunctuation, "{") {
		return nil
	}

	lit.Methods = p.parseInterfaceMethods()

	return lit
}

func (p *Parser) parseInterfaceMethods() []*ast.MethodDefinition {
	methods := []*ast.MethodDefinition{}

	p.nextToken()

	for !p.curTokenIs(lexer.TokenPunctuation) || p.curToken.Value != "}" {
		if p.curToken.Type == "EOF" {
			break
		}

		method := &ast.MethodDefinition{}
		// Method Name

		// Check canonical for 'func' (action) and skip it
		canonical := p.curToken.Value
		if p.dict != nil {
			if val, found := p.dict.Lookup(canonical); found {
				canonical = val
			}
		}

		if canonical == "func" || canonical == "action" {
			p.nextToken()
		}

		if !p.curTokenIs(lexer.TokenIdentifier) && !p.curTokenIs(lexer.TokenKeyword) {
			p.nextToken()
			continue
		}
		method.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}

		if !p.expectPeek(lexer.TokenPunctuation, "(") {
			return nil
		}

		method.Parameters = p.parseMethodParameters()

		// Return Type (Optional)
		// Check if next token is not punctuation that ends the line/block
		if !p.peekTokenIs(lexer.TokenPunctuation) && !p.peekTokenIs(lexer.TokenOperator) {
			p.nextToken()
			method.ReturnType = p.parseExpression(LOWEST)
		}

		methods = append(methods, method)

		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ";" {
			p.nextToken()
		}
		p.nextToken()
	}
	return methods
}

func (p *Parser) parseMethodParameters() []*ast.FieldDefinition {
	params := []*ast.FieldDefinition{}

	if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ")" {
		p.nextToken()
		return params
	}

	p.nextToken()

	for {
		param := &ast.FieldDefinition{}
		param.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}
		p.nextToken()
		param.Type = p.parseExpression(LOWEST)

		params = append(params, param)

		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == ")" {
			break
		}

		if p.peekTokenIs(lexer.TokenPunctuation) && p.peekToken.Value == "," {
			p.nextToken()
			p.nextToken()
			continue
		}

		// error?
		break
	}

	if !p.expectPeek(lexer.TokenPunctuation, ")") {
		return nil
	}

	return params
}

func (p *Parser) parseFunctionLiteral() *ast.FunctionLiteral {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(lexer.TokenPunctuation, "(") {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	// Parse return type if present (not {)
	if !p.peekTokenIs(lexer.TokenPunctuation) || p.peekToken.Value != "{" {
		lit.ReturnType = p.parseType() // Consumes the type
		// Double check if parseType consumes the token. Yes.
	}

	if !p.expectPeek(lexer.TokenPunctuation, "{") {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}
