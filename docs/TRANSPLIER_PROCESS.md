# The Transpiler Fix & Improvement Workflow

This document outlines a solid, repeatable process for fixing and improving the Singlish transpiler. Use this blueprint whenever you encounter a syntax error, a missing language feature, or a failing script.

## 1. Identify Failing Scripts (Regression Testing)

Always start by checking the current health of the transpiler against the existing examples to pinpoint what needs fixing.

* Run the test suite: `go build -o singlish main.go && go run run_all.go`.
* Take note of which scripts fail and their error messages (e.g., `expected next token to be punctuation...`).

## 2. Isolate & Debug the Issue

Take one failing script (e.g., `examples/123_new_feature.singlish`) and figure out *where* the transpiler is breaking. The pipeline is **Lexer → Parser → Codegen**.

* **Lexer Error:** Usually means a character or symbol isn't recognized (rare, but happens with new operators).
* **Parser Error:** The most common. The token stream doesn't match the expected grammar (e.g., `parseType`, `parseExpression`).
* **Codegen Error:** The script parses fine, but the Go compiler complains about the output generated in the temporary `.go` file.
* **Pro Tip:** Create a debug script (like `debug_ast.go`) to print out the Lexer diagnostics and the Parser errors directly for the specific failing file without the noise of the main CLI.

## 3. Implement the Fix across the Pipeline

Depending on what you found in step 2, you'll need to update one or more of these core packages:

* **`pkg/ast/ast.go` (The Blueprint):** If it's a completely new language construct, define the new AST Node struct and its `String()` method here.
* **`pkg/parser/parser.go` (The Logic):** Update the recursive descent parsing functions. If it's an expression, you might need to register it in `registerPrefix`/`registerInfix` or add it to the `precedence` map. If it's a statement, update the `parseStatement` branching logic.
* **`pkg/codegen/codegen.go` (The Output):** Add a `.visit...()` function for your new AST node so that the generator knows exactly what Go syntax to write out to the temporary file. Also, ensure your new node is handled in `.inspect()`.
* **`pkg/dictionaries/defaults.go` (The Vocabulary):** If you are adding new Singlish keyword mappings, map them here.

## 4. Verify the Local Fix

Rebuild the compiler and run the specific script that was failing:

```bash
go build -o singlish main.go
./singlish run examples/123_new_feature.singlish
```

If you encounter obscure Go compiler errors at this stage, use `./singlish transpile examples/...` to inspect the generated `.go` file directly.

## 5. Ensure No Regressions & Commit

Once the individual script works, run the full suite again to ensure your changes didn't break existing features.

* `go run run_all.go`
* If everything passes (or fails only for expected blocks like context timeouts), you are ready to stage, commit, and push!
