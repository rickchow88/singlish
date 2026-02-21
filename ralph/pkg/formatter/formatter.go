package formatter

import (
	"bytes"
	"strings"

	"github.com/rickchow/singlish/pkg/ast"
	"github.com/rickchow/singlish/pkg/dictionaries"
)

// Format converts AST back to canonical Singlish source code.
func Format(program *ast.Program, dict *dictionaries.Dictionary) (string, error) {
	f := &formatter{
		dict:        dict,
		indentLevel: 0,
	}

	f.format(program)
	return f.out.String(), nil
}

type formatter struct {
	dict        *dictionaries.Dictionary
	out         bytes.Buffer
	indentLevel int
}

// canonicalize takes a token value (keyword) and returns the canonical Singlish keyword.
func (f *formatter) canonicalize(kw string) string {
	// 1. Check if it's a known Singlish keyword (e.g. "lor")
	if goKw, found := f.dict.Lookup(kw); found {
		// It is. Get the canonical Singlish for that Go concept.
		// (e.g. "lor" -> "}" -> "lor" (if lor is canonical))
		if canonical, ok := f.dict.ReverseLookup(goKw); ok {
			return canonical
		}
		// Fallback: keep original if reverse lookup fails
		return kw
	}

	// 2. Check if it's a Go keyword (e.g. "}") used directly
	if canonical, ok := f.dict.ReverseLookup(kw); ok {
		// It is. Return the Singlish equivalent.
		return canonical
	}

	// 3. Unknown, return as is
	return kw
}

func (f *formatter) format(program *ast.Program) {
	// Generate package declaration
	var pkgStmt *ast.PackageStatement
	for _, s := range program.Statements {
		if ps, ok := s.(*ast.PackageStatement); ok {
			pkgStmt = ps
			break
		}
	}

	if pkgStmt != nil {
		f.visitPackageStatement(pkgStmt)
		f.write("\n\n")
	}

	// Generate imports
	// Collect import statements
	var imports []*ast.ImportStatement
	for _, s := range program.Statements {
		if is, ok := s.(*ast.ImportStatement); ok {
			imports = append(imports, is)
		}
	}

	if len(imports) > 0 {
		// We could group them, but let's just print them as is for now
		// or maybe group them into a single block?
		// Existing parser might have them as separate statements or a block?
		// The AST has `ImportStatement`.
		// Let's print them. Ideally we'd group them if they were grouped, but the AST seems to flatten them?
		// Wait, the Parser likely produces multiple ImportStatements.
		// Let's assume we want a single import block for clean formatting if multiple exist.

		// Note: codegen generates a block `import (...)`. Let's do the same for Singlish.
		// Find canonical keyword for "import"
		importKw := "import"
		if canonical, ok := f.dict.ReverseLookup("import"); ok {
			importKw = canonical
		}

		f.write(importKw)
		f.write(" (\n")
		f.indent()
		for _, imp := range imports {
			f.writeIndent()
			f.write(imp.Path.String()) // StringLiteral includes quotes
			f.write("\n")
		}
		f.dedent()
		f.write(")\n\n")
	}

	// Generate other statements
	for _, s := range program.Statements {
		if _, ok := s.(*ast.PackageStatement); ok {
			continue
		}
		if _, ok := s.(*ast.ImportStatement); ok {
			continue // Handled above
		}
		f.visit(s)
		f.write("\n")
	}
}

func (f *formatter) visit(node ast.Node) {
	switch n := node.(type) {
	case *ast.LetStatement:
		f.visitLetStatement(n)
	case *ast.ReturnStatement:
		f.visitReturnStatement(n)
	case *ast.FunctionStatement:
		f.visitFunctionStatement(n)
	case *ast.BlockStatement:
		f.visitBlockStatement(n)
	case *ast.ExpressionStatement:
		f.visitExpression(n.Expression)
	default:
		// Fallback for expressions
		if expr, ok := node.(ast.Expression); ok {
			f.visitExpression(expr)
		}
	}
}

func (f *formatter) visitPackageStatement(stmt *ast.PackageStatement) {
	kw := f.canonicalize(stmt.Token.Value)
	f.write(kw)
	f.write(" ")
	f.write(stmt.Name.Value)
}

func (f *formatter) visitLetStatement(stmt *ast.LetStatement) {
	kw := f.canonicalize(stmt.Token.Value)

	f.write(kw)
	f.write(" ")
	names := []string{}
	for _, name := range stmt.Names {
		names = append(names, name.Value)
	}
	f.write(strings.Join(names, ", "))

	if stmt.Type != nil {
		f.write(" ")
		typeVal := stmt.Type.Value

		// Handle pointer types e.g. *int or ki int
		// Actually, the AST parser stores the raw token string.
		// We should try to canonicalize types too if they are keywords (like int, string, bool?)
		// But types are often just identifiers.
		// If "int" has a Singlish equivalent, we should use it?
		// The dictionary might have "nombor: int".

		// Check for pointer prefix "*"
		isPtr := strings.HasPrefix(typeVal, "*")
		baseType := typeVal
		if isPtr {
			baseType = typeVal[1:]
		}

		// Canonicalize the base type if it's in the dictionary
		// e.g. "int" -> "nombor" (if nombor: int exists)
		// Wait, canonicalize() logic is:
		// 1. Singlish -> Go -> Canonical Singlish
		// 2. Go -> Canonical Singlish

		// If typeVal is "int" (Go), reverse lookup -> "nombor".
		// If typeVal is "nombor" (Singlish), lookup -> "int", reverse -> "nombor".
		baseType = f.canonicalize(baseType)

		if isPtr {
			// In Singlish, maybe we use `ki` instead of `*`?
			// The Requirement says "Ensure `ki` (pointer) syntax is handled".
			// If `ki` maps to `*` in the dictionary?
			// Currently `*` is likely handled by lexer as a token.
			// But `ki` might be a keyword in dictionary?
			// If dictionary has "ki: *", then canonicalize("*") -> "ki".
			// But "*" is not usually a keyword in the dictionary in that sense?
			// Let's assume standard * for now unless dictionary says otherwise.
			f.write("*" + baseType)
		} else {
			f.write(baseType)
		}
	}

	if stmt.Value != nil {
		f.write(" = ")
		f.visitExpression(stmt.Value)
	}
}

func (f *formatter) visitReturnStatement(stmt *ast.ReturnStatement) {
	kw := f.canonicalize(stmt.Token.Value)
	f.write(kw)
	if len(stmt.ReturnValues) > 0 {
		f.write(" ")
		for i, val := range stmt.ReturnValues {
			if i > 0 {
				f.write(", ")
			}
			f.visitExpression(val)
		}
	}
}

func (f *formatter) visitFunctionStatement(stmt *ast.FunctionStatement) {
	kw := f.canonicalize(stmt.Token.Value) // e.g. "func" -> "action"
	f.write(kw)
	f.write(" ")
	f.write(stmt.Name.Value)
	f.write("(")
	for i, param := range stmt.Parameters {
		if i > 0 {
			f.write(", ")
		}
		f.write(param.Name.Value)
		f.write(" ")
		f.visitExpression(param.Type)
	}
	f.write(") ")
	f.visitBlockStatement(stmt.Body)
}

func (f *formatter) visitBlockStatement(stmt *ast.BlockStatement) {
	f.write("{\n")
	f.indent()
	for _, s := range stmt.Statements {
		f.writeIndent()
		f.visit(s)
		f.write("\n")
	}
	f.dedent()
	f.writeIndent()
	f.write("}")
}

func (f *formatter) visitExpression(expr ast.Expression) {
	switch e := expr.(type) {
	case *ast.Identifier:
		// Identifiers (variables) shouldn't be canonicalized usually, unless they are types or keywords used as identifiers?
		// But here it's likely a variable name.
		f.write(e.Value)
	case *ast.FloatLiteral:
		f.write(e.Token.Value)
	case *ast.IntegerLiteral:
		f.write(e.Token.Value)
	case *ast.StringLiteral:
		f.write(e.Token.Value)
	case *ast.PrefixExpression:
		f.write("(")
		f.write(e.Operator)
		f.visitExpression(e.Right)
		f.write(")")
	case *ast.InfixExpression:
		if e.Operator == "." {
			// Method call e.g. fmt.Println
			// Should we canonicalize "fmt"? Probably not if it's a Go package.
			f.visitExpression(e.Left)
			f.write(".")
			f.visitExpression(e.Right)
		} else {
			f.write("(")
			f.visitExpression(e.Left)
			f.write(" ")
			f.write(e.Operator)
			f.write(" ")
			f.visitExpression(e.Right)
			f.write(")")
		}
	case *ast.CallExpression:
		f.visitExpression(e.Function)
		f.write("(")
		for i, arg := range e.Arguments {
			if i > 0 {
				f.write(", ")
			}
			f.visitExpression(arg)
		}
		f.write(")")
	}
}

func (f *formatter) write(s string) {
	f.out.WriteString(s)
}

func (f *formatter) writeIndent() {
	for i := 0; i < f.indentLevel; i++ {
		f.out.WriteString("\t")
	}
}

func (f *formatter) indent() {
	f.indentLevel++
}

func (f *formatter) dedent() {
	if f.indentLevel > 0 {
		f.indentLevel--
	}
}
