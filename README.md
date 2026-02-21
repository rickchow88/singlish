# Singlish 🇸🇬🦁

> *"Eh, you write code in Go ah? So cheem for what? Use Singlish lah!"*

**Singlish** is the world's first — and almost certainly *last* — **Go-based programming language** built entirely on the sacred colloquialisms of the **Singapore Kopitiam**. It is a high-performance, production-ready, fully compiled language for developers who are absolutely **sibeh sian** of typing `package`, `func`, and `nil` like some kind of ang moh.

Why write `package main` when you can write `kampung main`? Why say `func` like you're some kind of textbook when you can say `action`, like the pasar malam uncle who hasn't stopped hustling since 1987? Every keyword in Singlish has been **lovingly ripped from the streets of Singapore**, seasoned with decades of hawker centre wisdom, and welded onto the Go runtime with the structural integrity of a HDB void deck renovation.

This is not a toy. This is not a joke language. Well — okay, it *is* a joke language. But it compiles to real Go, runs at real Go speed, and will make your CI/CD pipeline question its life choices in perfect Singlish.

---

## 🤔 The Origin Story (Very Legit One)

It started, as all great Singaporean inventions do, with someone complaining. The founder looked at Go's syntax, then at their teh-tarik, and thought: *"This code got no soul one. Where's the kampung spirit?"*

Weeks of questionable engineering decisions later, `package` became `kampung`, `func` became `action`, `nil` became `kosong` (empty, like your bank account after Orchard Road), and `true` became `can` — because in Singapore, `can` is the answer to literally everything.

The **Insult Engine** was added shortly after, because a polite compiler error is a wasted opportunity.

---

## 💡 Why Singlish? (Confirm Plus Chop Got Reasons)

- **🚀 Go Speed, Kopitiam Soul** — Every line of Singlish compiles **1:1 into optimised Go code**. You get goroutines, garbage collection, static typing, and all the good stuff — but you also get to write `loop` instead of `for`, `balek` instead of `return`, and `nasi` instead of `if`. Your binary will be fast. Your code will be shiok.

- **🏃 Concurrency, Singapore Style** — Need to run things in parallel? Just `chiong` it (rush it, like you're late for the 7:31 MRT). Pass data through a `lobang` (channel). The whole concurrency model maps so naturally to Singlish that goroutines feel less like threading primitives and more like dispatching uncles to do different tasks simultaneously at a busy zi char stall.

- **😤 The "Kena Scold" Compiler** — Most compilers give you polite, passive-aggressive error messages. Singlish does not. If your code is wrong, the compiler will tell you — *loudly*, in hall-monitor Singlish. Expect gems like:
  - *"Eh bodoh, what is this?"*
  - *"Catch no ball sia"*
  - *"Wake up your idea"*
  - *"Blur like sotong"*
  - *"Simi sai is this?"*

  This is not cruelty. This is **mentorship**, Singapore-style.

- **📦 Full Go Interop, No Drama** — `dapao` (takeaway/import) any Go standard library or third-party package directly into your Singlish project. Need `net/http`? `dapao "net/http"`. Need `sort`? `dapao "sort"`. The entire Go ecosystem is yours. You're basically just writing Go, but make it *garang*.

- **🧠 Keywords That Actually Make Sense** — `kosong` for `nil` (empty). `bolehtak` for `bool` (can or not?). `nombor` for `int` (number). `tar` for `string` (talk/speech). `buat` for `make` (Malay: to make). Every keyword was chosen so that a Singapore auntie who has never programmed in her life could, theoretically, read your code and have *some* idea what's going on. We have not tested this theory. We are too scared.

- **🇸🇬 Localized Semantics** — The entry point is `action boss()`. The package is your `kampung`. You `chiong` goroutines, pass data through `lobang`s (gaps/channels), and when things go wrong you `gabra` (panic). When they recover, you say `heng ah` (lucky). This is not just a language. This is a **lifestyle**.

- **143 Examples and Counting** — From Hello World to HTTP servers, goroutine worker pools, cryptography, sorting algorithms, and a fully working **Singapore lottery number picker** (`88_lucky88.singlish`). Every example is a love letter to the Little Red Dot.

---

## 🚫 Who Should NOT Use Singlish

- People who enjoy writing `package` and `func` unironically
- Developers who think compiler error messages should be "encouraging"
- Anyone who says "lah" fewer than 3 times per sentence in real life
- Fans of TypeScript *(you know what you did)*

Everyone else: welcome to the kampung. `can` one.

---

## 🛠 Prerequisites

Before you start, ensure you have the following installed:

- **Go** (v1.21 or later) — [Install Go](https://go.dev/doc/install)
- **Git**

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

The repository includes **142 examples** in the `examples/` directory. Here are some to get you started:

### 1. Hello World (The Classic)

A simple demonstration of `kampung`, `action boss`, and `gong`.

```bash
./singlish run examples/01_hello_world.singlish
```

### 2. FizzBuzz (Logic & Loops)

Shows off `loop`, `nasi`, and `den` logic.

```bash
./singlish run examples/02_fizzbuzz.singlish
```

### 3. Goroutines (Concurrency)

Watch how Singlish handles `chiong` and `lobang` for high-performance concurrency.

```bash
./singlish run examples/03_coffee_shop_queue.singlish
```

### 4. Run All Examples

Run all scripts in one shot and see the summary:

```bash
go build -o singlish main.go && go run run_all.go
```

### 5. Using AI agents (Antigravity, Codex, Gemini, Claude, etc.)

> Run all singlish scripts in `/examples` and fix errors, run 1 test at a time directly from the terminal.

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
| `kaki` | `interface` | Interface |
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

> **Note**: If your code contains errors, the compiler will provide feedback in the form of authentic Singlish diagnostics (e.g., *"Wake up your idea"*, *"Blur like sotong"*, *"Eh bodoh"*, *"Catch no ball sia"*, or *"Simi sai is this?"*).

> **Gotcha**: `menu`, `count`, `buat`, `upsize`, `buang`, `kwear`, `gong`, `dun`, `or`, `somemore` are **reserved keywords** — do not use them as variable or loop-variable names.

## 📂 Examples (142 total)

Explore all [142 examples here](/examples).

### Foundation (01–20)

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

### Standard Library (21–50)

1. [21_cmd_args.singlish](/examples/21_cmd_args.singlish)
2. [22_strings.singlish](/examples/22_strings.singlish)
3. [23_math.singlish](/examples/23_math.singlish)
4. [24_time.singlish](/examples/24_time.singlish)
5. [25_tickers.singlish](/examples/25_tickers.singlish)
6. [26_mutex.singlish](/examples/26_mutex.singlish)
7. [27_waitgroup.singlish](/examples/27_waitgroup.singlish)
8. [28_atomic.singlish](/examples/28_atomic.singlish)
9. [29_context_timeout.singlish](/examples/29_context_timeout.singlish)
10. [30_context_cancel.singlish](/examples/30_context_cancel.singlish)
11. [31_lucky_numbers.singlish](/examples/31_lucky_numbers.singlish)
12. [32_makan_order.singlish](/examples/32_makan_order.singlish)
13. [33_sorting.singlish](/examples/33_sorting.singlish)
14. [34_string_functions.singlish](/examples/34_string_functions.singlish)
15. [35_regex.singlish](/examples/35_regex.singlish)
16. [36_url_parsing.singlish](/examples/36_url_parsing.singlish)
17. [37_sha1_hashes.singlish](/examples/37_sha1_hashes.singlish)
18. [38_base64_encoding.singlish](/examples/38_base64_encoding.singlish)
19. [39_random_numbers.singlish](/examples/39_random_numbers.singlish)
20. [40_number_parsing.singlish](/examples/40_number_parsing.singlish)
21. [41_xml_handling.singlish](/examples/41_xml_handling.singlish)
22. [42_custom_errors.singlish](/examples/42_custom_errors.singlish)
23. [43_environment_variables.singlish](/examples/43_environment_variables.singlish)
24. [44_subprocesses.singlish](/examples/44_subprocesses.singlish)
25. [45_signals.singlish](/examples/45_signals.singlish)
26. [46_exit.singlish](/examples/46_exit.singlish)
27. [47_embedding.singlish](/examples/47_embedding.singlish)
28. [48_method_sets.singlish](/examples/48_method_sets.singlish)
29. [49_interface_embedding.singlish](/examples/49_interface_embedding.singlish)
30. [50_select_with_default.singlish](/examples/50_select_with_default.singlish)

### Intermediate (51–80)

1. [51_closing_channels.singlish](/examples/51_closing_channels.singlish)
2. [52_timers.singlish](/examples/52_timers.singlish)
3. [53_type_assertion.singlish](/examples/53_type_assertion.singlish)
4. [54_closures.singlish](/examples/54_closures.singlish)
5. [55_fibonacci_closure.singlish](/examples/55_fibonacci_closure.singlish)
6. [56_bubble_sort.singlish](/examples/56_bubble_sort.singlish)
7. [57_binary_search.singlish](/examples/57_binary_search.singlish)
8. [58_named_returns.singlish](/examples/58_named_returns.singlish)
9. [59_multiple_return.singlish](/examples/59_multiple_return.singlish)
10. [60_error_wrapping.singlish](/examples/60_error_wrapping.singlish)
11. [61_iota_enum.singlish](/examples/61_iota_enum.singlish)
12. [62_string_builder.singlish](/examples/62_string_builder.singlish)
13. [63_sprintf_formatting.singlish](/examples/63_sprintf_formatting.singlish)
14. [64_bytes_buffer.singlish](/examples/64_bytes_buffer.singlish)
15. [65_buffered_channels.singlish](/examples/65_buffered_channels.singlish)
16. [66_worker_pool.singlish](/examples/66_worker_pool.singlish)
17. [67_pipeline.singlish](/examples/67_pipeline.singlish)
18. [68_rate_limiting.singlish](/examples/68_rate_limiting.singlish)
19. [69_sync_once.singlish](/examples/69_sync_once.singlish)
20. [70_sync_map.singlish](/examples/70_sync_map.singlish)
21. [71_rwmutex.singlish](/examples/71_rwmutex.singlish)
22. [72_stringer_interface.singlish](/examples/72_stringer_interface.singlish)
23. [73_linked_list.singlish](/examples/73_linked_list.singlish)
24. [74_stack.singlish](/examples/74_stack.singlish)
25. [75_queue.singlish](/examples/75_queue.singlish)
26. [76_higher_order_funcs.singlish](/examples/76_higher_order_funcs.singlish)
27. [77_composite_literals.singlish](/examples/77_composite_literals.singlish)
28. [78_goroutine_fanout.singlish](/examples/78_goroutine_fanout.singlish)
29. [79_init_function.singlish](/examples/79_init_function.singlish)
30. [80_slice_tricks.singlish](/examples/80_slice_tricks.singlish)

### Advanced (81–122)

1. [81_map_tricks.singlish](/examples/81_map_tricks.singlish)
2. [82_string_conversion.singlish](/examples/82_string_conversion.singlish)
3. [83_type_aliases.singlish](/examples/83_type_aliases.singlish)
4. [84_json_unmarshal.singlish](/examples/84_json_unmarshal.singlish)
5. [85_defer_order.singlish](/examples/85_defer_order.singlish)
6. [86_select_timeout.singlish](/examples/86_select_timeout.singlish)
7. [87_sync_pool.singlish](/examples/87_sync_pool.singlish)
8. [88_atomic_value.singlish](/examples/88_atomic_value.singlish)
9. [89_reflect_typeof.singlish](/examples/89_reflect_typeof.singlish)
10. [90_context_value.singlish](/examples/90_context_value.singlish)
11. [91_error_is_as.singlish](/examples/91_error_is_as.singlish)
12. [92_sort_custom.singlish](/examples/92_sort_custom.singlish)
13. [93_sync_cond.singlish](/examples/93_sync_cond.singlish)
14. [94_channel_range.singlish](/examples/94_channel_range.singlish)
15. [95_testing_mock.singlish](/examples/95_testing_mock.singlish)
16. [96_embedded_interfaces.singlish](/examples/96_embedded_interfaces.singlish)
17. [97_struct_tags.singlish](/examples/97_struct_tags.singlish)
18. [98_method_expressions.singlish](/examples/98_method_expressions.singlish)
19. [99_closures_state.singlish](/examples/99_closures_state.singlish)
20. [100_panic_defer.singlish](/examples/100_panic_defer.singlish)
21. [101_empty_interface.singlish](/examples/101_empty_interface.singlish)
22. [102_type_methods.singlish](/examples/102_type_methods.singlish)
23. [103_string_contains.singlish](/examples/103_string_contains.singlish)
24. [104_string_hasprefix.singlish](/examples/104_string_hasprefix.singlish)
25. [105_string_hassuffix.singlish](/examples/105_string_hassuffix.singlish)
26. [106_math_min.singlish](/examples/106_math_min.singlish)
27. [107_math_max.singlish](/examples/107_math_max.singlish)
28. [108_time_unix.singlish](/examples/108_time_unix.singlish)
29. [109_os_hostname.singlish](/examples/109_os_hostname.singlish)
30. [110_bytes_reader.singlish](/examples/110_bytes_reader.singlish)
31. [111_sort_float64.singlish](/examples/111_sort_float64.singlish)
32. [112_crypto_md5.singlish](/examples/112_crypto_md5.singlish)
33. [113_filepath_dir.singlish](/examples/113_filepath_dir.singlish)
34. [114_filepath_base.singlish](/examples/114_filepath_base.singlish)
35. [115_regexp_match.singlish](/examples/115_regexp_match.singlish)
36. [116_url_escape.singlish](/examples/116_url_escape.singlish)
37. [117_json_marshalindent.singlish](/examples/117_json_marshalindent.singlish)
38. [118_time_parse.singlish](/examples/118_time_parse.singlish)
39. [119_fmt_errorf.singlish](/examples/119_fmt_errorf.singlish)
40. [120_sync_waitgroup_loop.singlish](/examples/120_sync_waitgroup_loop.singlish)
41. [121_atomic_add.singlish](/examples/121_atomic_add.singlish)
42. [122_reflect_value.singlish](/examples/122_reflect_value.singlish)

### New: Singapore Edition (123–142)

1. [123_string_split.singlish](/examples/123_string_split.singlish) — `strings.Split` / `strings.Join`
2. [124_multiple_assign.singlish](/examples/124_multiple_assign.singlish) — Multiple return values & swap
3. [125_slice_of_structs.singlish](/examples/125_slice_of_structs.singlish) — Kopitiam menu board
4. [126_nested_maps.singlish](/examples/126_nested_maps.singlish) — Map inside a map (HDB seating)
5. [127_func_map.singlish](/examples/127_func_map.singlish) — Higher-order function + GST transform
6. [128_string_trim.singlish](/examples/128_string_trim.singlish) — `TrimSpace`, `TrimPrefix`, `TrimSuffix`
7. [129_goroutine_countdown.singlish](/examples/129_goroutine_countdown.singlish) — Goroutine countdown
8. [130_map_counting.singlish](/examples/130_map_counting.singlish) — Tally map (RC election)
9. [131_pointer_struct.singlish](/examples/131_pointer_struct.singlish) — Bank account with pointer receivers
10. [132_rune_handling.singlish](/examples/132_rune_handling.singlish) — Unicode/rune iteration
11. [133_type_switch.singlish](/examples/133_type_switch.singlish) — `see_how v.(type)` dispatch
12. [134_semaphore.singlish](/examples/134_semaphore.singlish) — Buffered channel semaphore
13. [135_strconv.singlish](/examples/135_strconv.singlish) — `Atoi`, `Itoa`, `ParseFloat`
14. [136_builder_pattern.singlish](/examples/136_builder_pattern.singlish) — Method chaining (kopi order)
15. [137_os_args.singlish](/examples/137_os_args.singlish) — `os.Args` CLI arguments
16. [138_func_dispatch.singlish](/examples/138_func_dispatch.singlish) — Results dispatch table
17. [139_memoization.singlish](/examples/139_memoization.singlish) — Cached Fibonacci
18. [140_interface_stack.singlish](/examples/140_interface_stack.singlish) — Generic stack with `interface{}`
19. [141_sort_slice.singlish](/examples/141_sort_slice.singlish) — `sort.Slice` custom comparator
20. [142_stringer_enum.singlish](/examples/142_stringer_enum.singlish) — MRT line enum with `String()`

## Test with

> go build -o singlish main.go && go run run_all.go

---

## 🤖 Development Loop (Ralph)

This project uses the **Ralph** agent loop for autonomous development. If you are a developer contributor:

1. Install dependencies: `npm install`.
2. Run a build iteration: `./bin/ralph build 1`.

---
*Made with ❤️ in Singapore.*
