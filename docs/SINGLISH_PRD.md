# Product Requirements Document (PRD): Project SINGLISH

**Date:** 2026-01-27
**Status:** Approved for Implementation
**Target Audience:** AI Developer / Transpiler Engineer

## 1. Executive Summary

**Project SINGLISH** is a systems programming language that serves as a culturally authentic "skin" over Go. It allows developers to write code using **Singaporean English (Singlish)** keywords and semantics, which are then transpiled into standard Go code for compilation.

The project is built as a **transpiler** (CLI tool) that:

1. Reads `.singlish` source files.
2. Tokenizes and parses Singlish syntax.
3. Translates it 1:1 into validated `.go` source code.
4. Executes the standard Go compiler (`go run` or `go build`).

## 2. Technical Foundation

- **Architecture:** Source-to-Source Transpiler (Singlish -> Go).
- **Implementation Language:** Go.
- **Output:** Standard Go Source Code.
- **File Extension:** `.singlish`
- **Error Handling Strategy:** "Kena Scold" (The compiler insults you when you fail).

## 3. Scope of Work

The goal is to build the `singlish` CLI tool. This tool handles the entire lifecycle from reading source code to producing a running binary.

### 3.1 Core Features

1. **Lexical Analysis**: A robust lexer to identify Singlish keywords (e.g., `kampung`, `lor`, `can`, `cannot`).
2. **Transpilation Engine**: A translation layer that maps Singlish tokens to Go tokens 1:1, preserving line numbers for debugging where possible.
3. **Insulting Error Handler**: A custom error reporter that wraps Go compiler errors with Singlish insults.
4. **Standard Library Interop**: Automatic aliasing of common standard library functions (like `fmt`).

## 4. Comprehensive Task List

### Phase 1: Project Initialization

- [ ] **Task 1.1**: Initialize Go Module `github.com/rickchow/singlish`.
- [ ] **Task 1.2**: scaffolding CLI structure (main.go, cmd/root.go, cmd/build.go, cmd/run.go).
- [ ] **Task 1.3**: Create `pkg/dictionaries` to hold the authoritative Singlish->Go keyword maps (loaded from `docs/SINGLISH_KEYWORDS.md`).

### Phase 2: The Transpiler Core (Lexer & Mapper)

- [ ] **Task 2.1**: **Implement Lexer**: Create `pkg/lexer` to read source code and emit tokens.
  - Support comments (`//`, `/* */`).
  - Support string literals (`"`, `` ` ``).
  - Identify identifiers vs keywords.
- [ ] **Task 2.2**: **Implement Keyword Mapper**: Create `pkg/mapper` to translating tokens.
  - **Structure**: `kampung`->`package`, `dapao`->`import`, `action`->`func`.
  - **Flow**: `nasi`->`if`, `den`->`else`, `loop`->`for`.
  - **Jumps**: `flykite`->`goto`, `tompang`->`fallthrough`.
  - **Concurrency**: `chiong`->`go`, `lobang`->`chan`, `pass`/`catch`->`<-`.
- [ ] **Task 2.3**: **Implement Type Mapper**:
  - `nombor`->`int`, `tar`->`string`, `bolehtak`->`bool`.
  - `cheem`->`complex128`, `zhi`->`rune`.
  - `ki` -> `*` (Pointer).

### Phase 3: Syntactic Sugar & StdLib

- [ ] **Task 3.1**: **Pointer Logic**: Ensure `ki int` translates correctly to `*int` and handle dereferencing if syntax differs.
- [ ] **Task 3.2**: **Value Translation**:
  - `can` -> `true`
  - `cannot` -> `false`
  - `kosong` -> `nil`
- [ ] **Task 3.3**: **Built-in Functions**:
  - Translate `gong(...)` -> `fmt.Println(...)` (or `fmt.Print` depending on context).
  - Translate `buat`->`make`, `kwear`->`close`, `buang`->`delete`.

### Phase 4: "Kena Scold" Error System

- [ ] **Task 4.1**: **Error Interceptor**: Capture stderr from the underlying `go build` command.
- [ ] **Task 4.2**: **Insult Engine**: Create a mechanism to inject insults before/after the actual Go error.
  - *Example*: "Eh bodoh, see line 5. Unknown token 'x'. Wake up please."

### Phase 5: Testing & Verification

- [ ] **Task 5.1**: **Unit Tests**: Test individual keyword translations (e.g., input `nasi` -> output `if`).
- [ ] **Task 5.2**: **Integration Tests**:
  - Transpile and run `examples/01_hello_world.singlish`.
  - Transpile and run `examples/02_fizzbuzz.singlish`.
  - Transpile and run `examples/11_pointers.singlish` (Verify `ki` logic).
- [ ] **Task 5.3**: **Chaos Test**: Feed invalid code to ensure the Insult Engine fires correctly.

## 5. Keyword Reference (Source of Truth)

| Singlish | Go | Singlish | Go | Singlish | Go |
|:---|:---|:---|:---|:---|:---|
| `kampung` | package | `nasi` | if | `got` | var |
| `dapao` | import | `den` | else | `confirm` | const |
| `action` | func | `say` | case | `nombor` | int |
| `boss` | main | `see_how` | switch | `tar` | string |
| `chiong` | go | `flykite` | goto | `ki` | * |
| `lobang` | chan | `tompang` | fallthrough | `zhi` | rune |
| `buat` | make | `upsize` | append | `gong` | fmt.Println |
| `kwear` | close | `buang` | delete | `kaki` | interface |

*(See `docs/SINGLISH_KEYWORDS.md` for the complete list)*

## 6. Credits & Inspiration

- **CURSED Language**: [github.com/ghuntley/cursed](https://github.com/ghuntley/cursed) - The esoteric language that started it all.
- **Go**: The rock-solid foundation this skin is built upon.

## Appendix A: SINGLISH_KEYWORDS.md Content

# PROPOSED KEYWORDS FOR SINGLISH (Based on CURSED)

This document proposes a comprehensive and culturally authentic set of keywords for the **SINGLISH** programming language, mapping directly from standard **CURSED** keywords.

## 1. Design Philosophy

- **Authenticity**: Keywords should be terms actually heard in Singaporean kopitiams, offices, and streets (e.g., `chiong`, `dapao`, `shiok`).
- **Intuition**: The Singlish term should map somewhat logically to the programming concept (e.g., `dapao` for `import` implies "bringing something back to use").
- **Humor**: Maintaining the "cursed" spirit with tongue-in-cheek Singlish adaptations.

## 2. Full Keyword Mapping

### 2.1 Top-Level Structure

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Package** | `package` | `vibe` | **`kampung`** | A package is a community/village. Code lives in a *kampung*. |
| **Import** | `import` | `yeet` | **`dapao`** | To take away / pack. You existing libraries to *dapao* them into your code. |
| **Function** | `func` | `slay` | **`action`** | "Eh, see him *action* only." Defining an action to perform. |
| **Main** | `main` | `main` | **`boss`** | The main boss function. (Optional, can stick to `main`). |

### 2.2 Variables & Constants

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Variable** | `var` | `sus` | **`got`** | "Got [name] [type]". implies existence. |
| **Constant** | `const` | `facts` | **`confirm`** | "Confirm plus chop", it won't change. |
| **Iota** | `iota` | `count` | **`auto`** | Automatic counter for constants. |
| **Type Def** | `type` | `be_like` | **`pattern`** | "More pattern than badminton". Defining the pattern of data. |

### 2.3 Control Flow

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **If** | `if` | `ready` | **`nasi`** | Teochew for "If". "Nasi you want...". |
| **Else** | `else` | `otherwise` | **`den`** | "Kalo this... *den* that". (Logic flow: If -> Then/Else). |
| **Select** | `select` | `vibez` | **`tikam`** | Random selection / gamble. Perfect for `select`. |
| **Switch** | `switch` | `vibe_check` | **`see_how`** | "See how things are", then decide. |
| **Case** | `case` | `mood` | **`say`** | "Let's *say* it's this...". |
| **Fallthrough** | `fallthrough` | `yeet_down` | **`tompang`** | "Hitch a ride" to the next case. |
| **Default** | `default` | `basic` | **`anyhow`** | "Anyhow do". The fallback option. |
| **Goto** | `goto` | `zoom` | **`flykite`** | "Flykite to". Skip/Jump to label. |

### 2.4 Loops

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **For** | `for` | `bestie` | **`loop`** | Simple. Can also use **`again`**. |
| **Range** | `range` | `flex` | **`all`** | "Take *all* items". |
| **Break** | `break` | `ghosted` | **`cabut`** | "Cabut liao" (Run away / escape). |
| **Continue** | `continue` | `simp` | **`go`** | "Go go go". Keep going. |

### 2.5 Functions & Returns

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Return** | `return` | `damn` | **`balek`** | Malay for "return/go back". "Balek kampung". |
| **Defer** | `defer` | `later` | **`nanti`** | Malay for "later". "Nanti do". |

### 2.6 Concurrency (The "Shiok" Part)

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Go** | `go` | `stan` | **`chiong`** | To rush/charge. "Cindy *chiong* this task". Implies async speed. |
| **Channel** | `chan` | `dm` | **`lobang`** | A gap/hole/opportunity. You pass things through the *lobang*. |
| **Send** | `<-` | `spill` | **`pass`** | "Pass to lobang". |
| **Receive** | `<-` | (implicit) | **`catch`** | "Catch from lobang". |

### 2.7 Values & Types

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **True** | `true` | `based` | **`can`** | Positive affirmation. |
| **False** | `false` | `cringe` | **`cannot`** | Negative affirmation. |
| **Nil** | `nil` | `nah` | **`kosong`** | Malay for "empty/zero". |
| **Bool** | `bool` | `lit` | **`bolehtak`** | "Can or not?". The boolean type. |
| **Int** | `int` | `normie` | **`nombor`** | Number. |
| **Int64** | `int64` | `thicc` | **`banyak`** | "Many/Much". Large number. |
| **Float** | `float64` | `meal` | **`point`** | Decimal point. |
| **Complex** | `complex128` | `galaxy_brain` | **`cheem`** | "Cheem" (Profound/Deep). Complex numbers. |
| **String** | `string` | `tea` | **`tar`** | Talk/Speech. |
| **Struct** | `struct` | `squad` | **`barang`** | "Things". A collection of data. |
| **Error** | `error` | `rly` | **`salah`** | "Wrong". Error type. |
| **Panic** | `panic` | `shook` | **`gabra`** | Panic/clumsy confusion. |
| **Pointer** | `*T` | `ඞ` | **`ki`** | Teochew for "point". Points to a value. |
| **Char** | `rune` | `sip` | **`zhi`** | Teochew for Letter/Character. |
| **Recover** | `recover` | `clutch` | **`heng`** | "Heng ah" (Lucky). Recovered from disaster. |

### 2.8 Data Structures & Interfaces

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Interface** | `interface` | `vibe_check` (n/a) | **`kaki`** | "Same kaki" (Same group/type). Defines a shared behavior. |
| **Map** | `map` | `tea_spill` | **`menu`** | Key-Value pairs, like a food menu. |
| **Connect** | `concat` | `link` | **`sambung`** | Join strings or lists. |

### 2.9 Built-in Functions & Operators

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Make/New** | `make` | `cook` | **`buat`** | Malay for "Make". |
| **Append** | `append` | `thicc` | **`upsize`** | "Upsize the meal" (Make it bigger/add to it). |
| **Delete** | `delete` | `yeet_out` | **`buang`** | Malay for "Throw away". |
| **Len** | `len` | `math` | **`count`** | Simple count. |
| **Close** | `close` | `ghost` | **`kwear`** | Teochew for "Close". |
| **Print** | `print` | `spill` | **`gong`** | Hokkien for "Talk". |
| **And** | `&&` | `and` | **`somemore`** | "This one... somemore that one". |
| **Or** | `\|\|` | `or` | **`or`** | (Standard English is fine). |
| **Not** | `!` | `not` | **`dun`** | "Dun do this". |
