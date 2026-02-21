# Singlish 🇸🇬

Singlish is a powerful and expressive programming language that brings the vibrant spirit of Singapore to the world of software development. Built as a Go-based transpiler, it allows you to write high-performance applications using familiar Singaporean colloquialisms.

## 🛠 Prerequisites

Before you start, ensure you have the following installed:

* **Go** (v1.21 or later) — [Install Go](https://go.dev/doc/install)
* **Git**

## 🚀 Installation & Setup

To install the Singlish CLI tool (`singlish`), follow these steps:

1. **Clone the repository**:

    ```bash
    git clone https://github.com/rickchow88/singlish.git
    cd singlish
    ```

2. **Build the CLI**:

    ```bash
    go build -o singlish main.go
    ```

3. **Verify the installation**:

    ```bash
    ./singlish --help
    ```

## 🏃 Running the Examples

The repository includes a variety of examples in the `examples/` directory.

To run an example (e.g., the classic "Hello World"):

```bash
./singlish run examples/hello.singlish
```

### Try these other examples

* **Math Operations**: `./singlish run examples/05_operators.singlish`
* **Goroutines**: `./singlish run examples/16_goroutines.singlish`
* **Channels**: `./singlish run examples/17_channels.singlish`
* **Timeout Context**: `./singlish run examples/29_context_timeout.singlish`

## 📖 Language Reference

| Singlish Keyword | Go Equivalent | Concept |
| :--- | :--- | :--- |
| `kampung` | `package` | Package definition |
| `dapao` | `import` | Import (Takeaway) |
| `action` | `func` | Function declaration |
| `boss` | `main` | Main entry point |
| `got` | `var` | Variable declaration |
| `confirm` | `const` | Constant declaration |
| `auto` | `iota` | Auto-incrementing constant |
| `pattern` | `type` | Type definition |
| `nasi` | `if` | If condition |
| `den` | `else` | Else condition |
| `tikam` | `select` | Select (Channel Ops) |
| `see_how` | `switch` | Switch statement |
| `say` | `case` | Case statement |
| `tompang` | `fallthrough` | Fallthrough statement |
| `anyhow` | `default` | Default case |
| `flykite` | `goto` | Goto statement |
| `loop` | `for` | For loop |
| `all` | `range` | Range iteration |
| `cabut` | `break` | Break loop |
| `go` | `continue` | Continue loop |
| `balek` | `return` | Return statement |
| `nanti` | `defer` | Defer execution |
| `chiong` | `go` | Spawn Goroutine |
| `lobang` | `chan` | Channel definition |
| `pass` | `<-` | Send to channel |
| `catch` | `<-` | Receive from channel |
| `can` | `true` | Boolean true |
| `cannot` | `false` | Boolean false |
| `kosong` | `nil` | Nil/Null |
| `bolehtak` | `bool` | Boolean type |
| `nombor` | `int` | Integer type |
| `banyak` | `int64` | 64-bit Integer |
| `point` | `float64` | Float type |
| `cheem` | `complex128` | Complex number |
| `tar` | `string` | String type |
| `barang` | `struct` | Struct definition |
| `salah` | `error` | Error type |
| `gabra` | `panic` | Panic |
| `ki` | `*` | Pointer |
| `zhi` | `rune` | Rune type |
| `heng` | `recover` | Recover from panic |
| `kaki` | `interface{}` | Interface |
| `menu` | `map` | Map definition |
| `buat` | `make` | Make built-in |
| `upsize` | `append` | Append built-in |
| `buang` | `delete` | Delete built-in |
| `count` | `len` | Length built-in |
| `kwear` | `close` | Close channel |
| `gong` | `fmt.Println` | Print to console |
| `somemore` | `&&` | Logical AND |
| `dun` | `!` | Logical NOT |
| `or` | `\|\|` | Logical OR |

> **Note**: If your code contains errors, the compiler will provide feedback in the form of authentic Singlish diagnostics (e.g., *"Wake up your idea"* or *"Blur like sotong"*).

## 📂 Examples (First 50)

Explore all [80+ examples here](/examples).

1. [01_hello_world.singlish](/examples/01_hello_world.singlish)
2. [02_fizzbuzz.singlish](/examples/02_fizzbuzz.singlish)
3. [03_coffee_shop_queue.singlish](/examples/03_coffee_shop_queue.singlish)
4. [04_structs_and_interfaces.singlish](/examples/04_structs_and_interfaces.singlish)
5. [05_error_handling.singlish](/examples/05_error_handling.singlish)
6. [06_arrays_slices.singlish](/examples/06_arrays_slices.singlish)
7. [07_maps.singlish](/examples/07_maps.singlish)
8. [08_defer.singlish](/examples/08_defer.singlish)
9. [09_switch_case.singlish](/examples/09_switch_case.singlish)
10. [10_panic_recover.singlish](/examples/10_panic_recover.singlish)
11. [11_pointers.singlish](/examples/11_pointers.singlish)
12. [12_constants.singlish](/examples/12_constants.singlish)
13. [13_variadic.singlish](/examples/13_variadic.singlish)
14. [14_recursion.singlish](/examples/14_recursion.singlish)
15. [15_file_write.singlish](/examples/15_file_write.singlish)
16. [16_file_read.singlish](/examples/16_file_read.singlish)
17. [17_json_encode.singlish](/examples/17_json_encode.singlish)
18. [18_json_decode.singlish](/examples/18_json_decode.singlish)
19. [19_http_server.singlish](/examples/19_http_server.singlish)
20. [20_http_client.singlish](/examples/20_http_client.singlish)
21. [21_cmd_args.singlish](/examples/21_cmd_args.singlish)
22. [22_strings.singlish](/examples/22_strings.singlish)
23. [23_math.singlish](/examples/23_math.singlish)
24. [24_time.singlish](/examples/24_time.singlish)
25. [25_tickers.singlish](/examples/25_tickers.singlish)
26. [26_mutex.singlish](/examples/26_mutex.singlish)
27. [27_waitgroup.singlish](/examples/27_waitgroup.singlish)
28. [28_atomic.singlish](/examples/28_atomic.singlish)
29. [29_context_timeout.singlish](/examples/29_context_timeout.singlish)
30. [30_context_cancel.singlish](/examples/30_context_cancel.singlish)
31. [31_lucky_numbers.singlish](/examples/31_lucky_numbers.singlish)
32. [32_makan_order.singlish](/examples/32_makan_order.singlish)
33. [33_sorting.singlish](/examples/33_sorting.singlish)
34. [34_string_functions.singlish](/examples/34_string_functions.singlish)
35. [35_regex.singlish](/examples/35_regex.singlish)
36. [36_url_parsing.singlish](/examples/36_url_parsing.singlish)
37. [37_sha1_hashes.singlish](/examples/37_sha1_hashes.singlish)
38. [38_base64_encoding.singlish](/examples/38_base64_encoding.singlish)
39. [39_random_numbers.singlish](/examples/39_random_numbers.singlish)
40. [40_number_parsing.singlish](/examples/40_number_parsing.singlish)
41. [41_xml_handling.singlish](/examples/41_xml_handling.singlish)
42. [42_custom_errors.singlish](/examples/42_custom_errors.singlish)
43. [43_environment_variables.singlish](/examples/43_environment_variables.singlish)
44. [44_subprocesses.singlish](/examples/44_subprocesses.singlish)
45. [45_signals.singlish](/examples/45_signals.singlish)
46. [46_exit.singlish](/examples/46_exit.singlish)
47. [47_embedding.singlish](/examples/47_embedding.singlish)
48. [48_method_sets.singlish](/examples/48_method_sets.singlish)
49. [49_interface_embedding.singlish](/examples/49_interface_embedding.singlish)
50. [50_select_with_default.singlish](/examples/50_select_with_default.singlish)

---

## 🤖 Development Loop (Ralph)

This project uses the **Ralph** agent loop for autonomous development. If you are a developer contributor:

1. Install dependencies: `npm install`.
2. Run a build iteration: `./bin/ralph build 1`.

---
*Made with ❤️ in Singapore.*
