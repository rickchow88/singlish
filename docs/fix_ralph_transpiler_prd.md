# PRD: Fix Ralph Transpiler and Parser Issues

## 1. Problem Description

Running the integration tests (`go test -v ./tests/integration`) revealed multiple failures in the Ralph transpiler. The failures range from critical segmentation faults to parsing errors for common Go/Singlish patterns.

### Summary of Failures

1. **Segmentation Fault (Critical)**: `TestExamples/06_arrays_slices.singlish` causes a panic in `codegen.go`.
2. **Parsing Errors**:
    - Loops with `range` (mapped to `all`?) fail to parse.
    - Multiple variable declarations (`got f, err = ...`) fail to parse.
    - Float literals (`2.3`) fail to parse as integers.
    - Switch statements and variadic parameters fail.
3. **Generated Code Syntax Errors**: Many examples fail with `syntax error: unexpected newline, expected type` or `function main is undeclared`, indicating the AST was not correctly built or the code generation produced invalid Go syntax.

## 2. Root Cause Analysis

### 2.1 Segmentation Fault in Codegen

- **Location**: `ralph/pkg/codegen/codegen.go:104` inside `inspect()` for `*ast.ForStatement`.
- **Cause**: The parser's `parseForStatement` (and likely others) returns `nil` when it encounters syntax it doesn't understand (e.g., `loop i, v = all ...`). However, because the return type is `*ast.ForStatement`, this `nil` pointer is wrapped in an `ast.Statement` interface, making it non-nil.
- **Result**: `parseBlockStatement` appends this "typed nil" to the statement list. When `codegen` iterates over statements and type-switches to `*ast.ForStatement`, it matches, but the instance `n` is `nil`. Accessing `n.Init` causes a panic.

### 2.2 Missing Range Loop Support

- **Issue**: `parseForStatement` only supports C-style `for` (`init; cond; post`) and `while`-style `for` (`cond`). It does not handle assignment expressions or `range` clauses found in `06_arrays_slices.singlish` (`loop i, drinks = all orders`).
- **Gap**: The AST and Parser need to support `RangeLoopStatement` or an expanded `ForStatement`.

### 2.3 Multiple Variable Declaration

- **Issue**: `got f, err = ...` (`var f, err = ...`) is used in `15_file_write.singlish`, but `ast.LetStatement` and `parseLetStatement` only support a single identifier (`Name`).
- **Gap**: The parser expects a type or assignment after the first identifier, failing when it encounters a comma.

### 2.4 Float Parsing

- **Issue**: `parseIntegerLiteral` uses `strconv.ParseInt`, which fails on `2.3`. There is no `FloatLiteral` support.

## 3. Proposed Changes

### 3.1 Fix Typed Nil Panic (Priority: Critical)

**File**: `ralph/pkg/parser/parser.go`

- **Action**: Modify `parseStatement` and other dispatch functions to explicitly check if the returned concrete type is `nil` before returning it as an interface.
- **Alternative**: In `parseBlockStatement`, check if `stmt` (interface) holds a nil value before appending.
- **Code Concept**:

    ```go
    case "for":
        s := p.parseForStatement()
        if s == nil { return nil }
        return s
    ```

### 3.2 Implement Range Loop Parsing

**File**: `ralph/pkg/parser/parser.go`, `ralph/pkg/ast/ast.go`

- **Action**: Update `ForStatement` AST or create `RangeStatement`.
- **Parser**: Detect syntax `for key, val = range X` (or `all X`).
- **Codegen**: Generate valid Go `for key, val := range X { ... }`.

### 3.3 Support Multiple Variable Declarations

**File**: `ralph/pkg/ast/ast.go`, `ralph/pkg/parser/parser.go`, `ralph/pkg/codegen/codegen.go`

- **AST**: Change `LetStatement.Name` to `Names []*Identifier`.
- **Parser**: `parseLetStatement` should parse a comma-separated list of identifiers.
- **Codegen**: `visitLetStatement` should write all names joined by commas.

### 3.4 Support Float Literals

**File**: `ralph/pkg/lexer/lexer.go` (if needed), `ralph/pkg/parser/parser.go`, `ralph/pkg/ast/ast.go`

- **AST**: Add `FloatLiteral` node.
- **Parser**: `parseIntegerLiteral` should be renamed or split. Tokenizer should distinguish floats or Parser should handle `.` in numbers.

### 3.5 Fix Main Function Detection

- Some tests fail with `main undeclared`. Verify that `kampung main` parses correctly to `package main` and that the file containing `action boss` (mapped to `main`) is properly processed.

## 4. Verification Plan

1. Apply the fix for the critical Segfault (3.1).
2. Run `go test -v ./tests/integration` and confirm `TestExamples/06` no longer panics (it may still fail compilation if range loop support is missing, but shouldn't crash).
3. Implement parsing features (3.2, 3.3, 3.4).
4. Run all integration tests and verify `PASS`.

## 5. Implementation Status (Updated 2026-01-28)

### Completed Fixes

1. **✅ Fix Typed Nil Panic (3.1)** - Each switch case in `parseStatement` now explicitly returns `nil` if the concrete type is `nil`.

2. **✅ Implement Range Loop Parsing (3.2)** - `ForStatement` AST updated with `IsRange`, `Key`, `Value`, and `Iterable` fields. Parser detects `loop key, val = all collection` syntax.

3. **✅ Support Multiple Variable Declarations (3.3)** - `LetStatement.Names` now holds `[]*Identifier`. Parser handles comma-separated identifiers.

4. **✅ Support Float Literals (3.4)** - `FloatLiteral` AST node added. `parseNumberLiteral` dispatches to appropriate parser.

5. **✅ Support Multiple Return Values** - `ReturnStatement.ReturnValues` now holds `[]Expression`. Parser handles comma-separated return values.

6. **✅ Support Function Return Types** - `FunctionStatement.ReturnType` field added. Parser and codegen handle single and grouped return types.

7. **✅ Add Switch Statement Support** - `SwitchStatement` and `CaseStatement` AST nodes added. Parser handles `see_how`/`switch` and `say`/`case`, `anyhow`/`default`. Added `switch` case to `parseStatement`.

8. **✅ Array Type Declarations** - `parseType` now handles `[n]type` syntax for fixed-size arrays.

9. **✅ Variadic Function Parameters** - `parseTypeExpression` handles `...type` syntax for variadic parameters.

10. **✅ Compound Assignment Operators** - Added `+=`, `-=`, `*=`, `/=`, etc. to precedence map and codegen.

11. **✅ Slice Expansion in Function Calls** - `parseCallArguments` handles `slice...` syntax for variadic expansion.

12. **✅ Empty Example Files** - Added content to `12_constants.singlish` and `21_cmd_args.singlish`.

13. **✅ Implement Select Statement** - Added `SelectStatement` and `SelectCase` AST. Implemented `parseSelectStatement` and `visitSelectStatement`. Fixed keyword-boundary issues in `parseReturnStatement`.

14. **✅ Fix Channel Send Syntax** - Added special case in `codegen` to handle `chan.pass(val)` and `chan.<- (val)` by translating to Go's `chan <- val` syntax.

### Test Results Summary (Updated Feb 2026)

- **Passing**: 82/82 examples (100% ✅, accounting for intentional cancellations/exits)
- **Failing**: 0 examples

Key passing tests demonstrate that core functionality works:

- Hello World, FizzBuzz, Structs/Interfaces
- Switch/Case and Select statements
- Arrays/Slices/Slice Tricks, Maps/SyncMap, Defer, Panic/Recover, Pointers
- Variadic functions, Recursion
- File I/O, JSON, HTTP Server/Client
- Strings, Math, Time, Context Timeout/Cancel
- Mutex, WaitGroup, Atomic operations
- Constants, Command-line arguments
- Deadlock-free channel operations
- Sorting, Custom Errors, Method Sets, Interface Embedding
