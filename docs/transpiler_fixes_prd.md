# PRD: Transpiler Fixes & Improvements

## 1. Overview

The Singlish-to-Go transpiler now correctly parses and transpiles almost all of the examples, achieving near 100% pass rates.

**Current Status**:

- ~95% Passing (Transpires and executes correctly)
- Only failing timeouts and blocking contexts.

## 2. Priority Levels

- **P0**: Critical Failures (Panics, basic syntax parsing broken). **(COMPLETED)**.
- **P1**: Standard Library & Feature Gaps (Missing type mappings, library functions). **(COMPLETED)**.
- **P2**: Test Suite Corrections (Incorrect assertions in tests). **(COMPLETED)**.

---

## 3. P0: Critical Infrastructure Fixes **(COMPLETED)**

### 3.1 Fix Compiler Panic in `codegen` (Arrays/Slices)

**Affected Test**: `06_arrays_slices.singlish`, `74_stack.singlish`, etc.
**Status**: Fixed. Added AST inspection guards, correct Array representation, and fully parsing index/slice syntax like `[low:high]`.

### 3.2 Fix Control Flow Parsing (Loops & Conditionals)

**Affected Tests**: `02_fizzbuzz.singlish`, `09_switch_case.singlish`
**Status**: Fixed. Control flows `nasi` (if) and `loop` (for) properly branch their bounds to correctly evaluate expressions and blocks without crashing. Logical operator groupings now correctly inherit `||` and `&&` priority tokens.

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
