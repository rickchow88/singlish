# PRD: Transpiler Fixes & Improvements

## 1. Overview

The Singlish-to-Go transpiler behaves correctly for simple programs but fails on complex control flow, specific data structures, and standard library usage. This PRD outlines the prioritized tasks required to achieve 100% pass rate on the `/docs/examples` integration suite.

**Current Status**:

- ~30% Passing (Basic output, basic concurrency)
- ~70% Failing (Control flow, Parser errors, Runtime panics, Output mismatches)

## 2. Priority Levels

- **P0**: Critical Failures (Panics, basic syntax parsing broken).
- **P1**: Standard Library & Feature Gaps (Missing type mappings, library functions).
- **P2**: Test Suite Corrections (Incorrect assertions in tests).

---

## 3. P0: Critical Infrastructure Fixes

### 3.1 Fix Compiler Panic in `codegen` (Arrays/Slices)

**Affected Test**: `06_arrays_slices.singlish`
**Symptom**: `panic: runtime error: invalid memory address or nil pointer dereference` in `codegen.go:104`.
**Root Cause**: The AST inspection phase (`g.inspect`) likely encounters a nil pointer wrapped in an interface (e.g., a `*ast.ForStatement` which is nil, or a child node which is nil).
**Requirement**:

- Add nil checks in `codegen.go` `inspect` and `visit` methods before accessing fields.
- Investigate `parser` to ensure it doesn't emit nil AST nodes.

### 3.2 Fix Control Flow Parsing (Loops & Conditionals)

**Affected Tests**: `02_fizzbuzz.singlish`, `09_switch_case.singlish`
**Symptom**: `transpilation failed: expected next token to be punctuation, got identifier instead`.
**Root Cause**:

- The parser expects standard Go syntax or strict token placement that Singlish keywords (`loop`, `nasi`, `den`) are not satisfying.
- `nasi` (if) and `loop` (for) parsing logic likely fails to handle the block start `{` correctly or the condition expression parsing is greedy.
**Requirement**:
- Debug `parser.ParseIfStatement` and `parser.ParseForStatement`.
- Ensure Singlish keywords `loop` and `nasi` correctly transition to parsing their conditions and block bodies.

---

## 4. P1: Feature & Library Implementation

### 4.1 Fix Standard Library Mappings

**Affected Tests**: `20_http_client.singlish`, `22_strings.singlish`, `23_math.singlish`, `24_time.singlish`
**Symptom**: Likely failing due to untranslated function names or missing imports.
**Requirement**:

- Update `dictionaries/stdlib.json` (or equivalent) to map Singlish keywords to Go standard library functions.
  - `strings.ToUpper` / `lowercase` etc.
  - `math.Max` etc.
  - `time.Now`, `time.Sleep`.
- Ensure standard libraries (like `net/http`) are automatically imported if detected, or explicitly `dapao`'d.

### 4.2 Fix Context Package Usage

**Affected Tests**: `29_context_timeout.singlish`, `30_context_cancel.singlish`
**Symptom**: Fails to compile or run.
**Requirement**:

- Ensure `context` package is correctly aliased or mapped.
- Handle context cancellation functions (which return a `cancel` func).

---

## 5. P2: Test Data Corrections

### 5.1 Fix `11_pointers.singlish` Assertion

**Symptom**: `Expected output to contain "100", got "60"`.
**Analysis**: The code performs `10 + 50 = 60`. The test expectation is outdated.
**Requirement**:

- Update `examples_test.go` to expect "60" for this test case.

---

## 6. Implementation Plan / Checklist

- [ ] **Fix P2**: Update `11_pointers` test expectation immediately to turn it Green.
- [ ] **Fix P0 (Panic)**: Add nil guards in `codegen.go`.
- [ ] **Fix P0 (Parser)**: Debug `02_fizzbuzz` parsing. Trace the token stream.
- [ ] **Fix P1**: Iteratively fix `strings`, `math`, `time`, and `http` examples by updating dictionary mappings.
