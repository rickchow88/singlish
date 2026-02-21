# Testing Plan for Singlish Transpiler

This document outlines the plan to test the Singlish transpiler using the examples provided in `/docs/examples`.

## Objective

Ensure that the Singlish transpiler correctly transpiles all valid example files to Go, and that the generated Go code executes successfully.

## Strategy

We will extend the existing integration test suite (`ralph/tests/integration/examples_test.go`) to automatically discover and test all `.singlish` files in the examples directory.

### Test Process

1. **Discovery**: Scan `docs/examples` for all `*.singlish` files.
2. **Transpilation & Execution**: For each file:
    - Invoke the `singlish` CLI to transpile and run the file (using the `run` command).
    - Verify that the process exits with code 0.
3. **Output Verification**:
    - For specific examples where the output is known and deterministic (e.g., `Hello World`, `FizzBuzz`), compare the standard output against the expected string.
    - For other examples, ensure no runtime errors occur.
4. **Exclusions**:
    - Skip files known to intentionally fail or block (e.g., `debug_catch.singlish` which causes a deadlock).

## automated Test Implementation

The file `ralph/tests/integration/examples_test.go` will be updated to:

- Use `filepath.Glob` to find all example files.
- Iterate through them and run dynamic subtests (`t.Run`).
- Maintain a map of `filename -> expectedOutput` for precise assertions.
- Maintain a set of `excludedFiles` for unstable tests.

## Running the Tests

Execute the tests using standard Go tools from the `ralph` directory:

```bash
cd ralph
go test -v ./tests/integration
```

## Current Status (As of Jan 2026)

The following examples are passing:

- 01_hello_world.singlish
- 03_coffee_shop_queue.singlish
- 04_structs_and_interfaces.singlish
- 08_defer.singlish
- 10_panic_recover.singlish
- 19_http_server.singlish
- 26_mutex.singlish
- 27_waitgroup.singlish
- 28_atomic.singlish
- debug_pass.singlish
- hello.singlish

Other examples are currently failing and require fixes in the transpiler.
