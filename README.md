# Singlish 🇸🇬

Singlish is a powerful and expressive programming language that brings the vibrant spirit of Singapore to the world of software development. Built as a Go-based transpiler, it allows you to write high-performance applications using familiar Singaporean colloquialisms.

## 🛠 Prerequisites

Before you start, ensure you have the following installed:

*   **Go** (v1.21 or later) — [Install Go](https://go.dev/doc/install)
*   **Git**

## 🚀 Installation & Setup

To install the Singlish CLI tool (`singlish`), follow these steps:

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/rickchow88/singlish.git
    cd singlish
    ```

2.  **Build the CLI**:
    ```bash
    cd ralph
    go build -o singlish main.go
    ```

3.  **Verify the installation**:
    ```bash
    ./singlish --help
    ```

## 🏃 Running the Examples

The repository includes a variety of examples in the `docs/examples/` directory.

To run an example (e.g., the classic "Hello World"):
```bash
./ralph/singlish run docs/examples/hello.singlish
```

### Try these other examples:
*   **Math Operations**: `./ralph/singlish run docs/examples/05_operators.singlish`
*   **Goroutines**: `./ralph/singlish run docs/examples/16_goroutines.singlish`
*   **Channels**: `./ralph/singlish run docs/examples/17_channels.singlish`
*   **Timeout Context**: `./ralph/singlish run docs/examples/29_context_timeout.singlish`

## 📖 Language Reference

| Singlish Keyword | Go Equivalent | Concept |
| :--- | :--- | :--- |
| `kampung` | `package` | Package definition |
| `dapao` | `import` | Import (Takeaway) |
| `action` | `func` | Function declaration |
| `gong` | `fmt.Println` | Print to console (Speak) |
| `boss()` | `main()` | Main entry point |
| `cho jiki` | `struct` | Struct definition |
| `steady` | `true` | Boolean true |
| `rabak` | `false` | Boolean false |

> **Note**: If your code contains errors, the compiler will provide feedback in the form of authentic Singlish diagnostics (e.g., *"Wake up your idea"* or *"Blur like sotong"*).

## 🤖 Development Loop (Ralph)

This project uses the **Ralph** agent loop for autonomous development. If you are a developer contributor:

1.  Navigate to the `ralph` directory.
2.  Install dependencies: `npm install`.
3.  Run a build iteration: `ralph build 1`.

---
*Made with ❤️ in Singapore.*
