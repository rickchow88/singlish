# PRD: Refactor Singlish Keyword 'main' to 'boss'

## 1. Overview

The user wants to rename the entry point function name from `main` to `boss` across the entire ecosystem. This change affects documentation, example code, and the core transpiler logic.

**Objective**: Replace `main` with `boss` as the reserved keyword for the entry function in Singlish. The transpiler will map `boss` back to `main` when generating Go code to maintain compatibility with the Go runtime.

## 2. Scope of Changes

### 2.1 Documentation (`/docs` & `/ralph/docs`)

- **Files**:
  - `docs/SINGLISH_KEYWORDS.md`
  - `ralph/docs/SINGLISH-User-manual.md`
  - `ralph/docs/SINGLISH_KEYWORDS.md`
- **Action**: Search for references to `main()` as the entry point and update them to `boss()`.
- **Note**: Ensure the distinction remains that `kampung main` (package main) typically stays as `kampung main` unless the user wants to rename the package too. **Assumption**: Only the function `main` becomes `boss`. The package declaration `package main` (Go) / `kampung main` (Singlish) usually remains required for executables, but we will check if `boss` implies a package change. For now, we assume `action boss()` replaces `action main()`.

### 2.2 Example Code (`/docs/examples` & `/ralph/docs/examples`)

- **Files**: All `*.singlish` files in `docs/examples`, `ralph/docs/examples`, and root `hello.singlish`.
- **Action**: Replace `action main()` with `action boss()`.

### 2.3 Transpiler Logic (`/ralph/pkg`)

- **Dictionary**: Update `dictionaries` (json or go map) to map `boss` -> `main`.
- **Parser/Codegen**:
  - If the transpiler hardcodes `main` detection, update it to look for `boss`.
  - Ensure that when `action boss()` is parsed, it is generated as `func main()` in the final Go output so the binary is executable.

## 3. Implementation Plan

### Phase 1: Keyword Definition & Docs

1. Update `docs/SINGLISH_KEYWORDS.md` to list `boss` as the entry point.
2. Update `ralph/docs/SINGLISH-User-manual.md`.

### Phase 2: Transpiler Update

1. **Modify Dictionary**: Add/Update entry `boss: "main"`.
2. **Verify AST/Codegen**: Check `codegen.go` or `transpiler.go`.
    - If `boss` is just a standard identifier that gets translated via dictionary, adding it to the dictionary might be sufficient.
    - **CRITICAL CHECK**: If `main` was previously a passthrough, strictly mapping `boss` -> `main` deals with the Go requirement.

### Phase 3: Code Migration

1. Iterate all `.singlish` files.
2. Regex replace `action main\s*\(` with `action boss(`.

### Phase 4: Validation

1. Run `go test ./tests/integration` to ensure `01_hello_world` and others still pass (compile and run).

## 4. Verification Steps

- [ ] Run `grep -r "action main" .` to ensure zero matches in `.singlish` files.
- [ ] Run `singlish run docs/examples/01_hello_world.singlish` and verify output.
