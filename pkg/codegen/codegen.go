package codegen

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/rickchow/singlish/pkg/ast"
	"github.com/rickchow/singlish/pkg/dictionaries"
)

// Generate converts AST to Go source code.
func Generate(program *ast.Program, dict *dictionaries.Dictionary) (string, error) {
	g := &generator{
		dict:        dict,
		imports:     make(map[string]struct{}),
		userImports: make(map[string]struct{}),
	}

	// First pass: collect explicit imports and detect implicit usage
	g.inspect(program)

	// Second pass: generate code
	g.generate(program)

	return g.out.String(), nil
}

type generator struct {
	dict        *dictionaries.Dictionary
	out         bytes.Buffer
	indentLevel int
	imports     map[string]struct{} // Implicitly required imports (e.g. "fmt")
	userImports map[string]struct{} // Explicit imports from source
}

func (g *generator) inspect(node ast.Node) {
	// Simple inspection to find usages of "fmt" and explicit imports
	switch n := node.(type) {
	case *ast.Program:
		if n == nil {
			return
		}
		for _, s := range n.Statements {
			g.inspect(s)
		}
	case *ast.ImportStatement:
		if n == nil {
			return
		}
		// Clean quotes
		path := strings.Trim(n.Path.Value, "\"`")
		g.userImports[path] = struct{}{}
	case *ast.InfixExpression:
		if n == nil {
			return
		}
		// Check for fmt.Method
		if n.Operator == "." {
			if left, ok := n.Left.(*ast.Identifier); ok && left.Value == "fmt" {
				g.imports["fmt"] = struct{}{}
			}
		}
		g.inspect(n.Left)
		g.inspect(n.Right)
	case *ast.CallExpression:
		if n == nil {
			return
		}
		g.inspect(n.Function)
		for _, arg := range n.Arguments {
			g.inspect(arg)
		}
	case *ast.LetStatement:
		if n == nil {
			return
		}
		g.inspect(n.Value)
	case *ast.ReturnStatement:
		if n == nil {
			return
		}
		for _, val := range n.ReturnValues {
			g.inspect(val)
		}
	case *ast.FunctionStatement:
		if n == nil {
			return
		}
		g.inspect(n.Body)
	case *ast.BlockStatement:
		if n == nil {
			return
		}
		for _, s := range n.Statements {
			g.inspect(s)
		}
	case *ast.ExpressionStatement:
		if n == nil {
			return
		}
		g.inspect(n.Expression)
	case *ast.PrefixExpression:
		if n == nil {
			return
		}
		g.inspect(n.Right)
	case *ast.Identifier:
		if n == nil {
			return
		}
		val := n.Value
		if translated, found := g.dict.Lookup(val); found {
			val = translated
		}
		if strings.HasPrefix(val, "fmt.") {
			g.imports["fmt"] = struct{}{}
		}
	case *ast.TypeStatement:
		if n == nil {
			return
		}
		g.inspect(n.Value)
	case *ast.StructLiteral:
		if n == nil {
			return
		}
		for _, f := range n.Fields {
			g.inspect(f.Type)
		}
	case *ast.InterfaceLiteral:
		if n == nil {
			return
		}
		for _, m := range n.Methods {
			g.inspect(m.ReturnType)
			for _, p := range m.Parameters {
				g.inspect(p.Type)
			}
		}
	case *ast.IfStatement:
		if n == nil {
			return
		}
		g.inspect(n.Condition)
		g.inspect(n.Consequence)
		if n.AlternativeStmt != nil {
			g.inspect(n.AlternativeStmt)
		}
	case *ast.ForStatement:
		if n == nil {
			return
		}
		if n.IsRange {
			g.inspect(n.Iterable)
		} else {
			if n.Init != nil {
				g.inspect(n.Init)
			}
			if n.Condition != nil {
				g.inspect(n.Condition)
			}
			if n.Post != nil {
				g.inspect(n.Post)
			}
		}
		g.inspect(n.Body)
	case *ast.IncDecStatement:
		if n == nil {
			return
		}
	case *ast.GoStatement:
		if n == nil {
			return
		}
		g.inspect(n.Call)
	case *ast.DeferStatement:
		if n == nil {
			return
		}
		g.inspect(n.Call)
	case *ast.SwitchStatement:
		if n == nil {
			return
		}
		if n.Expression != nil {
			g.inspect(n.Expression)
		}
		for _, c := range n.Cases {
			for _, e := range c.Expressions {
				g.inspect(e)
			}
			g.inspect(c.Body)
		}
	case *ast.SelectStatement:
		if n == nil {
			return
		}
		for _, c := range n.Cases {
			if c.Comm != nil {
				g.inspect(c.Comm)
			}
			if c.Body != nil {
				g.inspect(c.Body)
			}
		}
	case *ast.FloatLiteral:

		if n == nil {
			return
		}
	case *ast.IntegerLiteral:
		if n == nil {
			return
		}
	case *ast.TypeAssertionExpression:
		if n == nil {
			return
		}
		g.inspect(n.Left)
		g.inspect(n.Type)
	}
}

func (g *generator) generate(program *ast.Program) {
	// Generate package declaration
	// Find the package statement
	var pkgStmt *ast.PackageStatement
	for _, s := range program.Statements {
		if ps, ok := s.(*ast.PackageStatement); ok {
			pkgStmt = ps
			break
		}
	}

	if pkgStmt != nil {
		g.write("package ")
		g.write(pkgStmt.Name.Value)
		g.write("\n\n")
	} else {
		// Default package main? Or omit?
		// Let's assume input has package statement usually.
		// If not, we might want to default to main or let it fail.
		// For now, if missing, we don't write it (or write main?)
		g.write("package main\n\n")
	}

	// Generate imports
	g.generateImports()

	// Generate other statements
	for _, s := range program.Statements {
		if _, ok := s.(*ast.PackageStatement); ok {
			continue
		}
		if _, ok := s.(*ast.ImportStatement); ok {
			continue // Handled in generateImports
		}
		g.visit(s)
		g.write("\n")
	}
}

func (g *generator) generateImports() {
	allImports := make(map[string]struct{})
	for k := range g.userImports {
		allImports[k] = struct{}{}
	}
	for k := range g.imports {
		allImports[k] = struct{}{}
	}

	if len(allImports) == 0 {
		return
	}

	g.write("import (\n")
	g.indent()
	for imp := range allImports {
		g.writeIndent()
		g.write(fmt.Sprintf("%q\n", imp))
	}
	g.dedent()
	g.write(")\n\n")
}

func (g *generator) visit(node ast.Node) {
	if node == nil {
		return
	}
	switch n := node.(type) {
	case *ast.LetStatement:
		g.visitLetStatement(n)
	case *ast.ReturnStatement:
		g.visitReturnStatement(n)
	case *ast.FunctionStatement:
		g.visitFunctionStatement(n)
	case *ast.SwitchStatement:
		g.visitSwitchStatement(n)
	case *ast.SelectStatement:
		g.visitSelectStatement(n)
	case *ast.BlockStatement:

		g.visitBlockStatement(n)
	case *ast.ExpressionStatement:
		g.visitExpression(n.Expression)
		// Expression statements usually end with newline, added by loop
	case *ast.TypeStatement:
		g.visitTypeStatement(n)
	case *ast.IfStatement:
		g.visitIfStatement(n)
	case *ast.ForStatement:
		g.visitForStatement(n)
	case *ast.IncDecStatement:
		g.visitIncDecStatement(n)
	case *ast.GoStatement:
		g.visitGoStatement(n)
	case *ast.DeferStatement:
		g.visitDeferStatement(n)
	case *ast.IndexExpression:
		g.visitIndexExpression(n)
	case *ast.TypeAssertionExpression:
		g.visitExpression(n.Left)
		g.write(".(")
		g.visitExpression(n.Type)
		g.write(")")
	default:
		// Fallback for expressions
		if expr, ok := node.(ast.Expression); ok {
			g.visitExpression(expr)
		}
	}
}

func (g *generator) visitLetStatement(stmt *ast.LetStatement) {
	// Translate keyword: var/const/let
	kw := stmt.Token.Value
	if translated, found := g.dict.Lookup(kw); found {
		kw = translated
	} else {
		// Default mappings if dictionary misses them
		if kw == "let" {
			kw = "var"
		}
	}

	g.write(kw)
	g.write(" ")

	names := []string{}
	for _, name := range stmt.Names {
		names = append(names, name.Value)
	}
	g.write(strings.Join(names, ", "))

	if stmt.Type != nil {
		g.write(" ")
		typeVal := stmt.Type.Value
		isPtr := strings.HasPrefix(typeVal, "*")
		baseType := typeVal
		if isPtr {
			baseType = typeVal[1:]
		}

		if translated, found := g.dict.Lookup(baseType); found {
			baseType = translated
		}

		if isPtr {
			g.write("*" + baseType)
		} else {
			g.write(baseType)
		}
	}

	if stmt.Value != nil {
		g.write(" = ")
		g.visitExpression(stmt.Value)
	}
}

func (g *generator) visitReturnStatement(stmt *ast.ReturnStatement) {
	kw := "return"
	// Check dictionary just in case
	if translated, found := g.dict.Lookup("return"); found {
		kw = translated
	}

	g.write(kw)
	if len(stmt.ReturnValues) > 0 {
		g.write(" ")
		for i, val := range stmt.ReturnValues {
			if i > 0 {
				g.write(", ")
			}
			g.visitExpression(val)
		}
	}
}

func (g *generator) visitFunctionStatement(stmt *ast.FunctionStatement) {
	g.write("func ")
	if stmt.Receiver != nil {
		g.write("(")
		g.write(stmt.Receiver.Name.Value)
		g.write(" ")
		g.visitExpression(stmt.Receiver.Type)
		g.write(") ")
	}
	funcName := stmt.Name.Value
	if translated, found := g.dict.Lookup(funcName); found {
		funcName = translated
	}
	g.write(funcName)
	g.write("(")
	for i, param := range stmt.Parameters {
		if i > 0 {
			g.write(", ")
		}
		g.write(param.Name.Value)
		g.write(" ")
		g.visitExpression(param.Type)

	}
	g.write(")")

	if stmt.ReturnType != nil {
		g.write(" ")
		g.visitExpression(stmt.ReturnType)
	}

	g.write(" ")
	g.visitBlockStatement(stmt.Body)
}

func (g *generator) visitBlockStatement(stmt *ast.BlockStatement) {
	g.write("{\n")
	g.indent()
	for _, s := range stmt.Statements {
		g.writeIndent()
		g.visit(s)
		g.write("\n")
	}
	g.dedent()
	g.writeIndent()
	g.write("}")
}

func (g *generator) visitExpression(expr ast.Expression) {
	switch e := expr.(type) {
	case *ast.Identifier:
		val := e.Value

		// Handle pointer types in identifiers (hacky fix for parseType returning "*Type")
		isPtr := strings.HasPrefix(val, "*")
		baseVal := val
		if isPtr {
			baseVal = val[1:]
		}

		if translated, found := g.dict.Lookup(baseVal); found {
			baseVal = translated
		}

		if isPtr {
			g.write("*" + baseVal)
		} else {
			g.write(baseVal)
		}
	case *ast.FloatLiteral:
		g.write(e.Token.Value)
	case *ast.IntegerLiteral:
		g.write(e.Token.Value)
	case *ast.StringLiteral:
		g.write(e.Token.Value) // Keep quotes
	case *ast.PrefixExpression:
		g.write("(")
		g.write(e.Operator)
		g.visitExpression(e.Right)
		g.write(")")
	case *ast.InfixExpression:
		if e.Operator == "." {
			// Handle channel send: channel.<- (value) or channel.pass(value)
			// This happens when 'pass' (which is mapped to '<-') is used as a method call: chan.pass(val)
			// The right side is parsed as a PrefixExpression with '<-' operator.
			if pref, ok := e.Right.(*ast.PrefixExpression); ok && pref.Operator == "<-" {
				g.visitExpression(e.Left)
				g.write(" <- ")
				// If right side is a call, the arguments are the value: chan.pass(val)
				if call, ok := pref.Right.(*ast.CallExpression); ok {
					if len(call.Arguments) > 0 {
						g.visitExpression(call.Arguments[0])
					}
				} else {
					g.visitExpression(pref.Right)
				}
				return
			}
			g.visitExpression(e.Left)
			g.write(".")
			g.visitExpression(e.Right)

		} else if e.Operator == ":=" || e.Operator == "=" ||
			e.Operator == "+=" || e.Operator == "-=" ||
			e.Operator == "*=" || e.Operator == "/=" ||
			e.Operator == "%=" || e.Operator == "&=" ||
			e.Operator == "|=" || e.Operator == "^=" {
			g.visitExpression(e.Left)
			g.write(" ")
			g.write(e.Operator)
			g.write(" ")
			g.visitExpression(e.Right)
		} else if e.Operator == "<-" {
			// Send statement c <- v
			g.visitExpression(e.Left)
			g.write(" <- ")
			g.visitExpression(e.Right)
		} else {
			g.write("(")
			g.visitExpression(e.Left)
			g.write(" ")
			g.write(e.Operator)
			g.write(" ")
			g.visitExpression(e.Right)
			g.write(")")
		}

	case *ast.CallExpression:
		g.visitExpression(e.Function)
		g.write("(")
		for i, arg := range e.Arguments {
			if i > 0 {
				g.write(", ")
			}
			g.visitExpression(arg)
		}
		g.write(")")

	case *ast.StructLiteral:
		g.visitStructLiteral(e)
	case *ast.InterfaceLiteral:
		g.visitInterfaceLiteral(e)
	case *ast.IncDecStatement:
		g.visitIncDecStatement(e)
	case *ast.CompositeLiteral:
		g.visitExpression(e.Type)
		g.write("{")
		for i, el := range e.Elements {
			if i > 0 {
				g.write(", ")
			}
			g.visitExpression(el)
		}
		g.write("}")
	case *ast.KeyValueExpression:
		g.visitExpression(e.Key)
		g.write(": ")
		g.visitExpression(e.Value)
	case *ast.FunctionLiteral:
		g.visitFunctionLiteral(e)
	case *ast.IndexExpression:
		g.visitIndexExpression(e)
	case *ast.TypeAssertionExpression:
		g.visitExpression(e.Left)
		g.write(".(")
		g.visitExpression(e.Type)
		g.write(")")
	}
}

func (g *generator) visitFunctionLiteral(lit *ast.FunctionLiteral) {
	g.write("func")
	g.write("(")
	for i, param := range lit.Parameters {
		if i > 0 {
			g.write(", ")
		}
		g.write(param.Name.Value)
		g.write(" ")
		g.visitExpression(param.Type)
	}
	g.write(") ")

	if lit.ReturnType != nil {
		g.visitExpression(lit.ReturnType)
		g.write(" ")
	}

	g.visitBlockStatement(lit.Body)
}

func (g *generator) write(s string) {
	g.out.WriteString(s)
}

func (g *generator) writeIndent() {
	for i := 0; i < g.indentLevel; i++ {
		g.out.WriteString("\t")
	}
}

func (g *generator) indent() {
	g.indentLevel++
}

func (g *generator) dedent() {
	if g.indentLevel > 0 {
		g.indentLevel--
	}
}

func (g *generator) visitTypeStatement(stmt *ast.TypeStatement) {
	g.write("type ")
	g.write(stmt.Name.Value)
	g.write(" ")
	g.visitExpression(stmt.Value)
}

func (g *generator) visitStructLiteral(expr *ast.StructLiteral) {
	g.write("struct {\n")
	g.indent()
	for _, field := range expr.Fields {
		g.writeIndent()
		g.write(field.Name.Value)
		g.write(" ")
		g.visitExpression(field.Type)
		if field.Tag != nil {
			g.write(" ")
			g.write(field.Tag.Value)
		}
		g.write("\n")
	}
	g.dedent()
	g.writeIndent()
	g.write("}")
}

func (g *generator) visitInterfaceLiteral(expr *ast.InterfaceLiteral) {
	g.write("interface {\n")
	g.indent()
	for _, method := range expr.Methods {
		g.writeIndent()
		g.write(method.Name.Value)
		g.write("(")
		for i, param := range method.Parameters {
			if i > 0 {
				g.write(", ")
			}
			g.write(param.Name.Value)
			g.write(" ")
			g.visitExpression(param.Type)
		}
		g.write(")")
		if method.ReturnType != nil {
			g.write(" ")
			g.visitExpression(method.ReturnType)
		}
		g.write("\n")
	}
	g.dedent()
	g.writeIndent()
	g.write("}")
}

func (g *generator) visitIfStatement(stmt *ast.IfStatement) {
	g.write("if ")
	g.visitExpression(stmt.Condition)
	g.write(" ")
	g.visitBlockStatement(stmt.Consequence)

	if stmt.AlternativeStmt != nil {
		g.write(" else ")
		g.visit(stmt.AlternativeStmt)
	}
}

func (g *generator) visitForStatement(stmt *ast.ForStatement) {
	g.write("for ")

	if stmt.IsRange {
		if stmt.Key != nil {
			g.write(stmt.Key.Value)
			if stmt.Value != nil {
				g.write(", " + stmt.Value.Value)
			}
			g.write(" := range ")
		} else {
			// for range iterable
			g.write("range ")
		}
		g.visitExpression(stmt.Iterable)
	} else {
		// Handle Init
		if stmt.Init != nil {
			// Go requires ShortVarDecl (:=) or simple stmt in init clause.
			// "var i = 0" is not allowed.
			if let, ok := stmt.Init.(*ast.LetStatement); ok {
				names := []string{}
				for _, name := range let.Names {
					names = append(names, name.Value)
				}
				g.write(strings.Join(names, ", "))
				g.write(" := ")
				g.visitExpression(let.Value)
			} else {
				g.visit(stmt.Init)
			}
			g.write("; ")
		} else if stmt.Post != nil {
			g.write("; ")
		}

		// Handle Condition
		if stmt.Condition != nil {
			g.visitExpression(stmt.Condition)
		}

		// Handle Post
		if stmt.Post != nil {
			g.write("; ")
			g.visit(stmt.Post)
		}
	}

	g.write(" ")
	g.visitBlockStatement(stmt.Body)
}

func (g *generator) visitIncDecStatement(stmt *ast.IncDecStatement) {
	g.visitExpression(stmt.Left)
	g.write(stmt.Operator)
}

func (g *generator) visitGoStatement(stmt *ast.GoStatement) {
	g.write("go ")
	g.visitExpression(stmt.Call)
}

func (g *generator) visitDeferStatement(stmt *ast.DeferStatement) {
	g.write("defer ")
	g.visitExpression(stmt.Call)
}

func (g *generator) visitIndexExpression(exp *ast.IndexExpression) {
	g.visitExpression(exp.Left)
	g.write("[")
	g.visitExpression(exp.Index)
	g.write("]")
}

func (g *generator) visitSwitchStatement(stmt *ast.SwitchStatement) {
	g.write("switch ")
	if stmt.Expression != nil {
		g.visitExpression(stmt.Expression)
		g.write(" ")
	}
	g.write("{\n")
	g.indent()
	for _, c := range stmt.Cases {
		g.writeIndent()
		if c.Default {
			g.write("default:\n")
		} else {
			g.write("case ")
			for i, e := range c.Expressions {
				if i > 0 {
					g.write(", ")
				}
				g.visitExpression(e)
			}
			g.write(":\n")
		}
		g.indent()
		for _, s := range c.Body.Statements {
			g.writeIndent()
			g.visit(s)
			g.write("\n")
		}
		g.dedent()
	}
	g.dedent()
	g.writeIndent()
	g.write("}")
}

func (g *generator) visitSelectStatement(stmt *ast.SelectStatement) {
	g.write("select {\n")
	g.indent()
	for _, c := range stmt.Cases {
		g.writeIndent()
		if c.Default {
			g.write("default:\n")
		} else {
			g.write("case ")
			if c.Comm != nil {
				g.visit(c.Comm)
			}
			g.write(":\n")
		}
		g.indent()
		if c.Body != nil {
			for _, s := range c.Body.Statements {
				g.writeIndent()
				g.visit(s)
				g.write("\n")
			}
		}
		g.dedent()
	}
	g.dedent()
	g.writeIndent()
	g.write("}")
}
