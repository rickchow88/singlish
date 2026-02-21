package ast

import (
	"bytes"
	"strings"

	"github.com/rickchow/singlish/pkg/lexer"
)

// Node is the base interface for all nodes in the AST.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents a statement node.
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression node.
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of the AST.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// Identifier represents an identifier (variable name, function name, etc.).
type Identifier struct {
	Token lexer.Token // the lexer.TokenIdentifier token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Value }
func (i *Identifier) String() string       { return i.Value }

// IntegerLiteral represents an integer literal.
type IntegerLiteral struct {
	Token lexer.Token // the lexer.TokenNumber token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Value }
func (il *IntegerLiteral) String() string       { return il.Token.Value }

// FloatLiteral represents a float literal.
type FloatLiteral struct {
	Token lexer.Token // the lexer.TokenNumber token
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Value }
func (fl *FloatLiteral) String() string       { return fl.Token.Value }

// StringLiteral represents a string literal.

// StringLiteral represents a string literal.
type StringLiteral struct {
	Token lexer.Token // the lexer.TokenString token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Value }
func (sl *StringLiteral) String() string       { return sl.Token.Value }

// PrefixExpression represents a prefix operator expression (e.g. -5, !true).
type PrefixExpression struct {
	Token    lexer.Token // The prefix token, e.g. ! or -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Value }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression represents an infix operator expression (e.g. 5 + 5).
type InfixExpression struct {
	Token    lexer.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Value }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

// IndexExpression represents an index expression (e.g. array[i]).
type IndexExpression struct {
	Token lexer.Token // The '[' token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Value }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}

// CallExpression represents a function call.
type CallExpression struct {
	Token     lexer.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Value }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

// TypeStatement represents a type definition (type Name Type).
type TypeStatement struct {
	Token lexer.Token // the 'type' token
	Name  *Identifier
	Value Expression
}

func (ts *TypeStatement) statementNode()       {}
func (ts *TypeStatement) TokenLiteral() string { return ts.Token.Value }
func (ts *TypeStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ts.TokenLiteral() + " ")
	out.WriteString(ts.Name.String())
	out.WriteString(" ")
	if ts.Value != nil {
		out.WriteString(ts.Value.String())
	}
	return out.String()
}

// FieldDefinition represents a field in a struct.
type FieldDefinition struct {
	Name *Identifier
	Type Expression // Identifier or other type expression
	Tag  *StringLiteral
}

func (fd *FieldDefinition) String() string {
	var out bytes.Buffer
	out.WriteString(fd.Name.String())
	out.WriteString(" ")
	out.WriteString(fd.Type.String())
	if fd.Tag != nil {
		out.WriteString(" ")
		out.WriteString(fd.Tag.String())
	}
	return out.String()
}

// StructLiteral represents a struct definition.
type StructLiteral struct {
	Token  lexer.Token // the 'struct' token
	Fields []*FieldDefinition
}

func (sl *StructLiteral) expressionNode()      {}
func (sl *StructLiteral) TokenLiteral() string { return sl.Token.Value }
func (sl *StructLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(sl.TokenLiteral())
	out.WriteString(" {")
	for _, f := range sl.Fields {
		out.WriteString(" ")
		out.WriteString(f.String())
		out.WriteString(";")
	}
	out.WriteString(" }")
	return out.String()
}

// MethodDefinition represents a method in an interface.
type MethodDefinition struct {
	Name       *Identifier
	Parameters []*FieldDefinition // treating params as named fields for simplicity
	ReturnType Expression
}

func (md *MethodDefinition) String() string {
	var out bytes.Buffer
	out.WriteString(md.Name.String())
	out.WriteString("(")
	params := []string{}
	for _, p := range md.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	if md.ReturnType != nil {
		out.WriteString(" ")
		out.WriteString(md.ReturnType.String())
	}
	return out.String()
}

// InterfaceLiteral represents an interface definition.
type InterfaceLiteral struct {
	Token   lexer.Token // the 'interface' token
	Methods []*MethodDefinition
}

func (il *InterfaceLiteral) expressionNode()      {}
func (il *InterfaceLiteral) TokenLiteral() string { return il.Token.Value }
func (il *InterfaceLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(il.TokenLiteral())
	out.WriteString(" {")
	for _, m := range il.Methods {
		out.WriteString(" ")
		out.WriteString(m.String())
		out.WriteString(";")
	}
	out.WriteString(" }")
	return out.String()
}

// PackageStatement represents a package declaration.
type PackageStatement struct {
	Token lexer.Token // the lexer.TokenKeyword or Identifier token
	Name  *Identifier
}

func (ps *PackageStatement) statementNode()       {}
func (ps *PackageStatement) TokenLiteral() string { return ps.Token.Value }
func (ps *PackageStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ps.TokenLiteral() + " ")
	out.WriteString(ps.Name.String())
	return out.String()
}

// ImportStatement represents an import declaration.
type ImportStatement struct {
	Token lexer.Token // the lexer.TokenKeyword or Identifier token
	Path  *StringLiteral
}

func (is *ImportStatement) statementNode()       {}
func (is *ImportStatement) TokenLiteral() string { return is.Token.Value }
func (is *ImportStatement) String() string {
	var out bytes.Buffer
	out.WriteString(is.TokenLiteral() + " ")
	out.WriteString(is.Path.String())
	return out.String()
}

// LetStatement represents a variable declaration (let/var/const).
type LetStatement struct {
	Token lexer.Token // the token (let, var, const)
	Names []*Identifier
	Type  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Value }
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")

	names := []string{}
	for _, n := range ls.Names {
		names = append(names, n.String())
	}
	out.WriteString(strings.Join(names, ", "))

	if ls.Type != nil {
		out.WriteString(" ")
		out.WriteString(ls.Type.String())
	}
	if ls.Value != nil {
		out.WriteString(" = ")
		out.WriteString(ls.Value.String())
	}
	return out.String()
}

// ReturnStatement represents a return statement.
type ReturnStatement struct {
	Token        lexer.Token // the 'return' token
	ReturnValues []Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Value }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if len(rs.ReturnValues) > 0 {
		vals := []string{}
		for _, v := range rs.ReturnValues {
			vals = append(vals, v.String())
		}
		out.WriteString(strings.Join(vals, ", "))
	}
	return out.String()
}

// ExpressionStatement represents a statement that is just an expression.
type ExpressionStatement struct {
	Token      lexer.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Value }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// BlockStatement represents a block of statements (e.g. function body, if block).
type BlockStatement struct {
	Token      lexer.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Value }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// FunctionStatement represents a function definition.
type FunctionStatement struct {
	Token      lexer.Token // the 'func' token
	Name       *Identifier
	Receiver   *FieldDefinition // for methods: func (r Receiver) Name...
	Parameters []*FieldDefinition
	ReturnType Expression // Function return type (can be single Type or Grouped)
	Body       *BlockStatement
}

func (fs *FunctionStatement) statementNode()       {}
func (fs *FunctionStatement) TokenLiteral() string { return fs.Token.Value }
func (fs *FunctionStatement) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fs.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fs.TokenLiteral() + " ")
	if fs.Receiver != nil {
		out.WriteString("(")
		out.WriteString(fs.Receiver.String())
		out.WriteString(") ")
	}
	out.WriteString(fs.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	if fs.ReturnType != nil {
		out.WriteString(fs.ReturnType.String() + " ")
	}
	out.WriteString(fs.Body.String())
	return out.String()
}

// FunctionLiteral represents an anonymous function definition.
type FunctionLiteral struct {
	Token      lexer.Token // The 'func' token
	Parameters []*FieldDefinition
	Body       *BlockStatement
	ReturnType Expression
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Value }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	if fl.ReturnType != nil {
		out.WriteString(fl.ReturnType.String())
		out.WriteString(" ")
	}
	out.WriteString(fl.Body.String())
	return out.String()
}

// IfStatement represents an if/else control flow.
type IfStatement struct {
	Token       lexer.Token // the 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement // else block, optional
	// We might need to handle "else if" chains.
	// In Go AST, Else is a Statement (often an IfStatement for else-if, or BlockStatement for else).
	// Let's stick to BlockStatement for simple else, or Statement to allow recursion?
	// To allow "else if", Alternative should be Statement.
	AlternativeStmt Statement
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Value }
func (is *IfStatement) String() string {
	var out bytes.Buffer
	out.WriteString("if ")
	out.WriteString(is.Condition.String())
	out.WriteString(" ")
	out.WriteString(is.Consequence.String())
	if is.AlternativeStmt != nil {
		out.WriteString(" else ")
		out.WriteString(is.AlternativeStmt.String())
	}
	return out.String()
}

// ForStatement represents a for loop.
// Supports:
// 1. for init; condition; post { body }
// 2. for condition { body } (while)
// 3. for key, val = all collection { body } (range)
type ForStatement struct {
	Token lexer.Token // the 'for' token

	// Standard for and while
	Init      Statement
	Condition Expression
	Post      Statement

	// Range support
	IsRange  bool
	Key      *Identifier
	Value    *Identifier
	Iterable Expression

	Body *BlockStatement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Value }
func (fs *ForStatement) String() string {
	var out bytes.Buffer
	out.WriteString("for ")
	if fs.IsRange {
		if fs.Key != nil {
			out.WriteString(fs.Key.String())
			if fs.Value != nil {
				out.WriteString(", " + fs.Value.String())
			}
			out.WriteString(" = range ")
		} else if fs.Value != nil {
			out.WriteString(fs.Value.String() + " = range ")
		}
		out.WriteString(fs.Iterable.String())
	} else {
		if fs.Init != nil {
			out.WriteString(fs.Init.String() + "; ")
		}
		if fs.Condition != nil {
			out.WriteString(fs.Condition.String())
		}
		if fs.Post != nil {
			out.WriteString("; " + fs.Post.String())
		}
	}
	out.WriteString(" ")
	out.WriteString(fs.Body.String())
	return out.String()
}

// IncDecStatement represents i++ or i--.
// These are statements in Go, not expressions.
type IncDecStatement struct {
	Token    lexer.Token // ++ or --
	Left     Expression
	Operator string
}

func (ids *IncDecStatement) statementNode()       {}
func (ids *IncDecStatement) expressionNode()      {}
func (ids *IncDecStatement) TokenLiteral() string { return ids.Token.Value }
func (ids *IncDecStatement) String() string {
	return ids.Left.String() + ids.Operator
}

// KeyValueExpression represents "Key: Value".
type KeyValueExpression struct {
	Token lexer.Token // the ':' token
	Key   Expression
	Value Expression
}

func (kve *KeyValueExpression) expressionNode()      {}
func (kve *KeyValueExpression) TokenLiteral() string { return kve.Token.Value }
func (kve *KeyValueExpression) String() string {
	var out bytes.Buffer
	out.WriteString(kve.Key.String())
	out.WriteString(": ")
	out.WriteString(kve.Value.String())
	return out.String()
}

// CompositeLiteral represents Type{ Elements }.
type CompositeLiteral struct {
	Token    lexer.Token  // the '{' token
	Type     Expression   // The type being instantiated (e.g. Person)
	Elements []Expression // The values (often KeyValueExpression)
}

// SwitchStatement represents a switch statement.
type SwitchStatement struct {
	Token      lexer.Token
	Expression Expression
	Cases      []*CaseStatement
}

func (ss *SwitchStatement) statementNode()       {}
func (ss *SwitchStatement) TokenLiteral() string { return ss.Token.Value }
func (ss *SwitchStatement) String() string {
	var out bytes.Buffer
	out.WriteString("switch ")
	if ss.Expression != nil {
		out.WriteString(ss.Expression.String())
	}
	out.WriteString(" {")
	for _, c := range ss.Cases {
		out.WriteString(c.String())
	}
	out.WriteString("}")
	return out.String()
}

// CaseStatement represents a case or default in a switch.
type CaseStatement struct {
	Token       lexer.Token
	Default     bool
	Expressions []Expression
	Body        *BlockStatement
}

func (cs *CaseStatement) String() string {
	var out bytes.Buffer
	if cs.Default {
		out.WriteString("default:")
	} else {
		out.WriteString("case ")
		exprs := []string{}
		for _, e := range cs.Expressions {
			exprs = append(exprs, e.String())
		}
		out.WriteString(strings.Join(exprs, ", "))
		out.WriteString(":")
	}
	out.WriteString(cs.Body.String())
	return out.String()
}

func (cl *CompositeLiteral) expressionNode()      {}
func (cl *CompositeLiteral) TokenLiteral() string { return cl.Token.Value }
func (cl *CompositeLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(cl.Type.String())
	out.WriteString("{")
	args := []string{}
	for _, el := range cl.Elements {
		args = append(args, el.String())
	}
	out.WriteString(strings.Join(args, ", "))
	out.WriteString("}")
	return out.String()
}

// GoStatement represents a go goroutine call.
type GoStatement struct {
	Token lexer.Token // the 'go' token
	Call  *CallExpression
}

func (gs *GoStatement) statementNode()       {}
func (gs *GoStatement) TokenLiteral() string { return gs.Token.Value }
func (gs *GoStatement) String() string {
	var out bytes.Buffer
	out.WriteString("go ")
	out.WriteString(gs.Call.String())
	return out.String()
}

// DeferStatement represents a defer statement.
type DeferStatement struct {
	Token lexer.Token // the 'defer' token
	Call  *CallExpression
}

func (ds *DeferStatement) statementNode()       {}
func (ds *DeferStatement) TokenLiteral() string { return ds.Token.Value }
func (ds *DeferStatement) String() string {
	var out bytes.Buffer
	out.WriteString("defer ")
	out.WriteString(ds.Call.String())
	return out.String()
}

// SelectStatement represents a select statement for channel operations.
type SelectStatement struct {
	Token lexer.Token // the 'select' token
	Cases []*SelectCase
}

func (ss *SelectStatement) statementNode()       {}
func (ss *SelectStatement) TokenLiteral() string { return ss.Token.Value }
func (ss *SelectStatement) String() string {
	var out bytes.Buffer
	out.WriteString("select {")
	for _, c := range ss.Cases {
		out.WriteString(c.String())
	}
	out.WriteString("}")
	return out.String()
}

// SelectCase represents a case in a select statement.
// Comm can be a send expression (c <- v) or a receive expression (<-c or x := <-c)
type SelectCase struct {
	Token   lexer.Token
	Default bool
	Comm    Statement // The communication clause (send or receive)
	Body    *BlockStatement
}

func (sc *SelectCase) statementNode()       {}
func (sc *SelectCase) TokenLiteral() string { return sc.Token.Value }
func (sc *SelectCase) String() string {
	var out bytes.Buffer
	if sc.Default {
		out.WriteString("default:")
	} else {
		out.WriteString("case ")
		if sc.Comm != nil {
			out.WriteString(sc.Comm.String())
		}
		out.WriteString(":")
	}
	if sc.Body != nil {
		out.WriteString(sc.Body.String())
	}
	return out.String()
}

type TypeAssertionExpression struct {
	Token lexer.Token
	Left  Expression
	Type  Expression
}

func (tae *TypeAssertionExpression) expressionNode()      {}
func (tae *TypeAssertionExpression) TokenLiteral() string { return tae.Token.Value }
func (tae *TypeAssertionExpression) String() string {
	var out bytes.Buffer
	out.WriteString(tae.Left.String())
	out.WriteString(".(")
	out.WriteString(tae.Type.String())
	out.WriteString(")")
	return out.String()
}
