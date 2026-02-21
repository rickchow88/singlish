# Singlish User Manual

## Table of Contents

- [Introduction](#introduction)
- [Installation & Prerequisites](#installation--prerequisites)
- [CLI Command Reference](#cli-command-reference)
- [Dictionary & Syntax Guide](#dictionary--syntax-guide)

## Introduction

**Singlish** is a delightful and humorous esoteric programming language wrapper for the Go programming language. It brings the colorful and expressive Singaporean English dialect (Singlish) into the world of software development.

With Singlish, you can write code using familiar colloquialisms like `kampung` (package), `dapao` (import), and `action` (func). The Singlish CLI tool then transpiles your `.sg` (Singlish) source files into standard Go code, which can be compiled and run using the standard Go toolchain.

This manual provides a comprehensive guide to installing the Singlish CLI, understanding the syntax and dictionary, and using the core commands to build and format your Singlish projects.

## Installation & Prerequisites

### Prerequisites

Before using Singlish, you must have the **Go programming language** installed on your system, as Singlish transpiles to Go.

- **Go (Golang)**: Download and install the latest version from [golang.org](https://go.dev/dl/).

### Installation

To install the Singlish CLI, follow these steps:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/rickchow/singlish.git
   cd singlish
   ```

2. **Build the CLI**:
   Run the following command in the project root to build the `singlish` binary:

   ```bash
   go build -o singlish main.go
   ```

   This will create an executable file named `singlish` (or `singlish.exe` on Windows) in your current directory.

### Adding to PATH

For convenience, you can add the binary to your system's PATH so you can run `singlish` from any directory.

#### Linux / macOS

Move the binary to `/usr/local/bin` (requires sudo):

```bash
sudo mv singlish /usr/local/bin/
```

Or add the current directory to your PATH in your shell configuration (e.g., `~/.bashrc`, `~/.zshrc`):

```bash
export PATH=$PATH:$(pwd)
```

#### Windows

Add the directory containing `singlish.exe` to your User PATH environment variable via System Properties.

## CLI Command Reference

The Singlish CLI supports several commands to manage your Singlish projects.

### Global Flags

The following flags can be used with any command:

- `--dictionary <path>`
  - Specifies the path to a custom dictionary file.
  - Default: `dictionary.txt` in the current directory.
  - Example: `singlish --dictionary=my_dict.txt build main.sg`

### Commands

#### `transpile`

Transpiles a Singlish source file into a temporary Go source file. This is useful if you want to inspect the generated Go code without compiling it immediately.

**Usage:**

```bash
singlish transpile <file>
```

**Example:**

```bash
singlish transpile main.sg
```

**Output:** Prints the path to the generated temporary Go file.

#### `build`

Transpiles the Singlish file and compiles it into an executable binary using the Go compiler.

**Usage:**

```bash
singlish build <file>
```

**Example:**

```bash
singlish build main.sg
```

**Output:** Creates an executable file (e.g., `main` or `main.exe`) in the current directory.

#### `fmt`

Formats a Singlish source file according to the canonical style. It updates the file in place.

**Usage:**

```bash
singlish fmt <file>
```

**Example:**

```bash
singlish fmt main.sg
```

#### `run`

Transpiles and immediately runs the Singlish file.

**Usage:**

```bash
singlish run <file>
```

**Example:**

```bash
singlish run main.sg
```

## Dictionary & Syntax Guide

### Dictionary Format

Singlish allows you to customize the keywords using a `dictionary.txt` file. The file format is simple: each line contains a Singlish keyword mapped to a Go keyword, separated by a colon (`:`).

- **Format:** `<Singlish Keyword> : <Go Keyword>`
- **Comments:** Lines starting with `#` or `//` are ignored.
- **Location:** The CLI looks for `dictionary.txt` in the current directory by default. You can also specify a custom path using the `--dictionary` flag.

**Example `dictionary.txt`:**

```text
# Custom Dictionary
kampung : package
dapao   : import
action  : func
lor     : ;  # Map 'lor' to semicolon
```

### Default Mappings

If no dictionary is provided, Singlish uses the default mappings. Here are the core defaults:

| Singlish | Go Keyword | Description |
| :--- | :--- | :--- |
| `kampung` | `package` | Package declaration |
| `dapao` | `import` | Import declaration |
| `action` | `func` | Function declaration |
| `nasi` | `if` | If statement |
| `pattern` | `type` | Type declaration |
| `susun` | `struct` | Struct declaration |
| `lobang` | `interface` | Interface declaration |
| `gong` | `fmt.Println` | Print to console |
| `nombor` | `int` | Integer type |
| `tar` | `string` | String type |
| `bolehtak` | `bool` | Boolean type |
| `ki` | `*` | Pointer |
| `buat` | `make` | Make built-in function |
| `buang` | `delete` | Delete built-in function |
| `upsize` | `append` | Append built-in function |
| `kwear` | `close` | Close built-in function |

### Syntax Cheat Sheet

Singlish syntax largely follows Go, but with localized keywords.

#### Variables

```go
// Singlish
var x nombor = 10
const y tar = "Hello"

// Go Equivalent
var x int = 10
const y string = "Hello"
```

#### Functions

```go
// Singlish
action boss() {
    gong("Hello World")
}

// Go Equivalent
func main() {
    fmt.Println("Hello World")
}
```

#### Control Flow (Loops & If)

Singlish uses standard Go syntax for loops (`for`), unless you map it in your dictionary.

```go
// Singlish
nasi x > 5 {
    gong("Big")
}

for i := 0; i < 5; i++ {
    gong(i)
}
```

#### Structs & Interfaces

```go
// Singlish
pattern Person susun {
    Name tar
    Age  nombor
}

pattern Greeter lobang {
    Greet()
}

// Go Equivalent
type Person struct {
    Name string
    Age  int
}

type Greeter interface {
    Greet()
}
```

## Examples

### Hello World

A classic "Hello World" program in Singlish.

**File:** `hello.sg`

```go
kampung main

action boss() {
    gong("Hello World")
}
```

**Run it:**

```bash
singlish run hello.sg
```

### Structs

Here is an example demonstrating how to define and use structs.

**File:** `people.sg`

```go
kampung main

// Define a struct
pattern Person susun {
    Name tar
    Age  nombor
}

// Define a function that takes a pointer to a struct
action SayHello(p ki Person) {
    gong("Hello, my name is " + p.Name)
}

action boss() {
    // Create a new Person
    var uncle Person
    
    // Set fields
    uncle.Name = "Uncle Roger"
    uncle.Age = 50

    // Call the function
    SayHello(&uncle)
}
```

## Troubleshooting

### Common Issues

#### `exec: "go": executable file not found in $PATH`

**Cause:** The Go programming language is not installed or not in your system's PATH.
**Solution:** Install Go from [golang.org](https://go.dev/dl/) and ensure the `go` command works in your terminal.

#### `unknown token: ...` or Syntax Errors

**Cause:** You might be using a word that isn't in the dictionary, or the syntax is incorrect.
**Solution:** Check your `dictionary.txt` (or default mappings) to ensure you are using the correct Singlish keywords. Run `singlish transpile <file>` to see the generated Go code and debug where the syntax might be broken.

#### `no such file or directory`

**Cause:** The file you are trying to run or build does not exist.
**Solution:** Verify the file path. Remember that Singlish looks for `dictionary.txt` in the current directory by default.

#### Compilation Errors (e.g., `undefined: fmt`)

**Cause:** The transpiled Go code has errors.
**Solution:**

- If you are using standard Singlish keywords like `gong`, the compiler handles `fmt` imports for you.
- If you are using custom mappings or raw Go code that requires imports, ensure you use `dapao` (import) to include the necessary packages.
