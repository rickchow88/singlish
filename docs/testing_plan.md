# Testing Plan for Singlish Transpiler

This document outlines the plan to test the Singlish transpiler using the examples provided in `/examples`.

## Objective

Ensure that the Singlish transpiler correctly transpiles all valid example files to Go, and that the generated Go code executes successfully.

## Strategy

We run all `.singlish` files via `run_all.go`, which uses `./singlish run <file>` with a 3-second per-file timeout.

### Test Process

1. **Discovery**: `run_all.go` scans `examples/` for all `*.singlish` files (alphabetical order).
2. **Transpilation & Execution**: For each file:
    - Invoke `./singlish run <file>`.
    - Verify that the process exits with code 0 within 3 seconds.
3. **Output Verification**: For examples where output is deterministic (e.g., Hello World, FizzBuzz), compare stdout against expected strings.
4. **Exclusions** (hardcoded in `run_all.go`):
    - `19_http_server.singlish` — blocks on port listen, no timeout
    - `debug_catch.singlish` — intentional deadlock demonstration

### Running All Tests

```bash
go build -o singlish main.go && go run run_all.go
```

### Running a Single Test

```bash
./singlish run examples/<file>.singlish
```

---

## Current Status (As of Feb 2026)

**142 example files** — all passing `run_all.go` except intentional exclusions.

### Intentional Exclusions / Known Behaviour

| File | Reason |
|------|--------|
| `19_http_server.singlish` | Blocks on HTTP listen — excluded from run_all |
| `debug_catch.singlish` | Deadlock by design |
| `46_exit.singlish` | Exits with code 3 by design (tested separately) |

### Previously Fixed Issues (Feb 2026)

| File | Issue | Fix |
|------|-------|-----|
| `45_signals.singlish` | `os.FindProcess` returns 2 values, only 1 captured | Added `proc, err` unpacking |
| `52_timers.singlish` | 2s timer exceeded 3s run_all timeout | Reduced to 500ms |
| `117_json_marshalindent.singlish` | Variable named `menu` clashed with `menu` keyword | Renamed to `dishes` |
| `121_atomic_add.singlish` | Variable named `count` clashed with `count` keyword | Renamed to `counter` |
| `124_multiple_assign.singlish` | `got a, b nombor = 10, 20` multi-var init unsupported | Split into separate `got` declarations |
| `125_slice_of_structs.singlish` | Variable named `menu` clashed with keyword | Renamed to `items` |
| `127_func_map.singlish` | `action(nombor) nombor` as func param type unsupported | Inlined as closure |
| `130_map_counting.singlish` | Loop var named `count` clashed with keyword | Renamed to `total` |
| `134_semaphore.singlish` | `struct{}` channel + inline `<-sem` closure unsupported | Use `nombor` channel + helper func |
| `138_func_dispatch.singlish` | `got a, b nombor = 12, 4` multi-var init unsupported | Split into separate `got` declarations |
| `139_memoization.singlish` | `nasi val, ok := m[n]; ok` multi-value if-init unsupported | Split to `got cached, ok = ...` then `nasi ok` |
| `140_interface_stack.singlish` | `nasi dun ok` parsing failure | Changed to `== cannot` |
| `142_stringer_enum.singlish` | `confirm (` grouped const block unsupported | Individual `confirm` with explicit values |

---

## Known Transpiler Limitations

| Limitation | Workaround |
|------------|-----------|
| Reserved keywords (`menu`, `count`, `buat`, `upsize`, `buang`, `kwear`, `gong`, `dun`, `or`, `somemore`) cannot be used as variable names | Use descriptive alternatives |
| `got a, b type = x, y` multi-var init with explicit type unsupported | Use separate `got` declarations |
| `action(T) T` as a function parameter or map-value type unsupported | Store as closure variable or pre-compute |
| `nasi val, ok := map[key]; ok` if-init form unsupported | Split into `got val, ok = map[key]` then `nasi ok` |
| `nasi dun boolVar` — `dun` before bare identifier fails | Use `== cannot` for negative bool checks |
| Bare `catch chan` after a `gong(...)` call merges into same line | Use `got _ = catch chan` form |
| `confirm (` grouped const block with `iota` unsupported | Use individual `confirm Name Type = value` |
| `struct{}` as channel element type (`lobang barang{}`) fails | Use `lobang nombor` as semaphore token type |
