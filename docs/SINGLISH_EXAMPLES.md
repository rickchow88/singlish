# Singlish Logic Examples

Here are **142 example programs** written in **SINGLISH** demonstrating the syntax and logic of the language. Each file is self-contained and runnable with `./singlish run <file>`.

---

## Foundation (01–10)

### 1. Hello World

File: [01_hello_world.singlish](../examples/01_hello_world.singlish)

- Basic printing (`gong`)
- Main function (`action boss`)

### 2. FizzBuzz

File: [02_fizzbuzz.singlish](../examples/02_fizzbuzz.singlish)

- Loop (`loop`)
- If/Else (`nasi`/`den`)
- Modulo arithmetic

### 3. Coffee Shop Queue (Concurrency)

File: [03_coffee_shop_queue.singlish](../examples/03_coffee_shop_queue.singlish)

- Goroutines (`chiong`)
- Channels (`lobang`)
- Sending/Receiving (`pass`/`catch`)

### 4. Structs & Interfaces (Kakis & Barang)

File: [04_structs_and_interfaces.singlish](../examples/04_structs_and_interfaces.singlish)

- Struct definitions (`pattern ... barang`)
- Interface definitions (`pattern ... kaki`)
- Methods (`action (p *Person) ...`)
- Booleans (`can`/`cannot`)

### 5. Error Handling (Salah & Heng)

File: [05_error_handling.singlish](../examples/05_error_handling.singlish)

- Multiple returns
- Error type (`salah`)
- Nil checks (`kosong`)

### 6. Arrays & Slices

File: [06_arrays_slices.singlish](../examples/06_arrays_slices.singlish)

- Fixed arrays
- Slices (`[]nombor`)
- Appending (`upsize`)
- Range loops (`loop ... all`)

### 7. Maps (Menu)

File: [07_maps.singlish](../examples/07_maps.singlish)

- Map creation (`buat(menu...)`)
- Key existence check
- Deleting keys (`buang`)

### 8. Defer (Nanti)

File: [08_defer.singlish](../examples/08_defer.singlish)

- Deferred execution (`nanti`)

### 9. Switch Case (See How)

File: [09_switch_case.singlish](../examples/09_switch_case.singlish)

- Switch (`see_how`)
- Case (`say`)
- Default (`anyhow`)

### 10. Panic & Recover (Gabra & Heng)

File: [10_panic_recover.singlish](../examples/10_panic_recover.singlish)

- Panic (`gabra`)
- Recover (`heng`)
- Named returns for recovery

---

## Standard Library (11–50)

| # | File | Topics |
|---|------|--------|
| 11 | [11_pointers.singlish](../examples/11_pointers.singlish) | Pointers (`ki`), pass by reference |
| 12 | [12_constants.singlish](../examples/12_constants.singlish) | `confirm`, `auto` (iota) |
| 13 | [13_variadic.singlish](../examples/13_variadic.singlish) | Variadic functions |
| 14 | [14_recursion.singlish](../examples/14_recursion.singlish) | Recursive functions |
| 15 | [15_file_write.singlish](../examples/15_file_write.singlish) | `os.WriteFile` |
| 16 | [16_file_read.singlish](../examples/16_file_read.singlish) | `os.ReadFile` |
| 17 | [17_json_encode.singlish](../examples/17_json_encode.singlish) | `json.Marshal` |
| 18 | [18_json_decode.singlish](../examples/18_json_decode.singlish) | `json.Unmarshal` |
| 19 | [19_http_server.singlish](../examples/19_http_server.singlish) | HTTP server (skip in run_all) |
| 20 | [20_http_client.singlish](../examples/20_http_client.singlish) | HTTP client |
| 21 | [21_cmd_args.singlish](../examples/21_cmd_args.singlish) | `os.Args` |
| 22 | [22_strings.singlish](../examples/22_strings.singlish) | `strings` package |
| 23 | [23_math.singlish](../examples/23_math.singlish) | `math` package |
| 24 | [24_time.singlish](../examples/24_time.singlish) | `time` package |
| 25 | [25_tickers.singlish](../examples/25_tickers.singlish) | `time.Ticker` |
| 26 | [26_mutex.singlish](../examples/26_mutex.singlish) | `sync.Mutex` |
| 27 | [27_waitgroup.singlish](../examples/27_waitgroup.singlish) | `sync.WaitGroup` |
| 28 | [28_atomic.singlish](../examples/28_atomic.singlish) | `sync/atomic` |
| 29 | [29_context_timeout.singlish](../examples/29_context_timeout.singlish) | `context.WithTimeout` |
| 30 | [30_context_cancel.singlish](../examples/30_context_cancel.singlish) | `context.WithCancel` |
| 31 | [31_lucky_numbers.singlish](../examples/31_lucky_numbers.singlish) | Random + Singlish flavour |
| 32 | [32_makan_order.singlish](../examples/32_makan_order.singlish) | Struct + methods |
| 33 | [33_sorting.singlish](../examples/33_sorting.singlish) | `sort.Ints`, `sort.Strings` |
| 34 | [34_string_functions.singlish](../examples/34_string_functions.singlish) | `strings.*` |
| 35 | [35_regex.singlish](../examples/35_regex.singlish) | `regexp` |
| 36 | [36_url_parsing.singlish](../examples/36_url_parsing.singlish) | `net/url` |
| 37 | [37_sha1_hashes.singlish](../examples/37_sha1_hashes.singlish) | `crypto/sha1` |
| 38 | [38_base64_encoding.singlish](../examples/38_base64_encoding.singlish) | `encoding/base64` |
| 39 | [39_random_numbers.singlish](../examples/39_random_numbers.singlish) | `math/rand` |
| 40 | [40_number_parsing.singlish](../examples/40_number_parsing.singlish) | `strconv` |
| 41 | [41_xml_handling.singlish](../examples/41_xml_handling.singlish) | `encoding/xml` |
| 42 | [42_custom_errors.singlish](../examples/42_custom_errors.singlish) | Custom error types |
| 43 | [43_environment_variables.singlish](../examples/43_environment_variables.singlish) | `os.Getenv` |
| 44 | [44_subprocesses.singlish](../examples/44_subprocesses.singlish) | `exec.Command` |
| 45 | [45_signals.singlish](../examples/45_signals.singlish) | `os/signal`, `syscall` |
| 46 | [46_exit.singlish](../examples/46_exit.singlish) | `os.Exit` |
| 47 | [47_embedding.singlish](../examples/47_embedding.singlish) | Struct embedding |
| 48 | [48_method_sets.singlish](../examples/48_method_sets.singlish) | Method sets |
| 49 | [49_interface_embedding.singlish](../examples/49_interface_embedding.singlish) | Interface embedding |
| 50 | [50_select_with_default.singlish](../examples/50_select_with_default.singlish) | `tikam` + `anyhow` |

---

## Intermediate (51–80)

| # | File | Topics |
|---|------|--------|
| 51 | [51_closing_channels.singlish](../examples/51_closing_channels.singlish) | `kwear`, range over channel |
| 52 | [52_timers.singlish](../examples/52_timers.singlish) | `time.NewTimer`, `Stop` |
| 53 | [53_type_assertion.singlish](../examples/53_type_assertion.singlish) | Type assertion |
| 54 | [54_closures.singlish](../examples/54_closures.singlish) | Closures |
| 55 | [55_fibonacci_closure.singlish](../examples/55_fibonacci_closure.singlish) | Fibonacci with closure |
| 56 | [56_bubble_sort.singlish](../examples/56_bubble_sort.singlish) | Bubble sort |
| 57 | [57_binary_search.singlish](../examples/57_binary_search.singlish) | Binary search |
| 58 | [58_named_returns.singlish](../examples/58_named_returns.singlish) | Named return values |
| 59 | [59_multiple_return.singlish](../examples/59_multiple_return.singlish) | Multiple return values |
| 60 | [60_error_wrapping.singlish](../examples/60_error_wrapping.singlish) | `fmt.Errorf` wrapping |
| 61 | [61_iota_enum.singlish](../examples/61_iota_enum.singlish) | Iota enum constants |
| 62 | [62_string_builder.singlish](../examples/62_string_builder.singlish) | `strings.Builder` |
| 63 | [63_sprintf_formatting.singlish](../examples/63_sprintf_formatting.singlish) | `fmt.Sprintf` |
| 64 | [64_bytes_buffer.singlish](../examples/64_bytes_buffer.singlish) | `bytes.Buffer` |
| 65 | [65_buffered_channels.singlish](../examples/65_buffered_channels.singlish) | Buffered channels |
| 66 | [66_worker_pool.singlish](../examples/66_worker_pool.singlish) | Goroutine worker pool |
| 67 | [67_pipeline.singlish](../examples/67_pipeline.singlish) | Channel pipeline |
| 68 | [68_rate_limiting.singlish](../examples/68_rate_limiting.singlish) | Rate limiting |
| 69 | [69_sync_once.singlish](../examples/69_sync_once.singlish) | `sync.Once` |
| 70 | [70_sync_map.singlish](../examples/70_sync_map.singlish) | `sync.Map` |
| 71 | [71_rwmutex.singlish](../examples/71_rwmutex.singlish) | `sync.RWMutex` |
| 72 | [72_stringer_interface.singlish](../examples/72_stringer_interface.singlish) | `fmt.Stringer` |
| 73 | [73_linked_list.singlish](../examples/73_linked_list.singlish) | Linked list |
| 74 | [74_stack.singlish](../examples/74_stack.singlish) | Stack data structure |
| 75 | [75_queue.singlish](../examples/75_queue.singlish) | Queue data structure |
| 76 | [76_higher_order_funcs.singlish](../examples/76_higher_order_funcs.singlish) | Higher-order functions |
| 77 | [77_composite_literals.singlish](../examples/77_composite_literals.singlish) | Nested composite literals |
| 78 | [78_goroutine_fanout.singlish](../examples/78_goroutine_fanout.singlish) | Goroutine fan-out |
| 79 | [79_init_function.singlish](../examples/79_init_function.singlish) | `init()` function |
| 80 | [80_slice_tricks.singlish](../examples/80_slice_tricks.singlish) | Slice manipulation tricks |

---

## Advanced (81–122)

| # | File | Topics |
|---|------|--------|
| 81 | [81_map_tricks.singlish](../examples/81_map_tricks.singlish) | Map manipulation |
| 82 | [82_string_conversion.singlish](../examples/82_string_conversion.singlish) | String/byte/rune conversion |
| 83 | [83_type_aliases.singlish](../examples/83_type_aliases.singlish) | Type aliases |
| 84 | [84_json_unmarshal.singlish](../examples/84_json_unmarshal.singlish) | `json.Unmarshal` into map |
| 85 | [85_defer_order.singlish](../examples/85_defer_order.singlish) | Defer LIFO order |
| 86 | [86_select_timeout.singlish](../examples/86_select_timeout.singlish) | Select with timeout |
| 87 | [87_sync_pool.singlish](../examples/87_sync_pool.singlish) | `sync.Pool` |
| 88 | [88_atomic_value.singlish](../examples/88_atomic_value.singlish) | `atomic.Value` |
| 89 | [89_reflect_typeof.singlish](../examples/89_reflect_typeof.singlish) | `reflect.TypeOf` |
| 90 | [90_context_value.singlish](../examples/90_context_value.singlish) | `context.WithValue` |
| 91 | [91_error_is_as.singlish](../examples/91_error_is_as.singlish) | `errors.Is` / `errors.As` |
| 92 | [92_sort_custom.singlish](../examples/92_sort_custom.singlish) | Custom sort interface |
| 93 | [93_sync_cond.singlish](../examples/93_sync_cond.singlish) | `sync.Cond` |
| 94 | [94_channel_range.singlish](../examples/94_channel_range.singlish) | Range over channel |
| 95 | [95_testing_mock.singlish](../examples/95_testing_mock.singlish) | Interface mock |
| 96 | [96_embedded_interfaces.singlish](../examples/96_embedded_interfaces.singlish) | Embedded interfaces |
| 97 | [97_struct_tags.singlish](../examples/97_struct_tags.singlish) | Struct tags + reflect |
| 98 | [98_method_expressions.singlish](../examples/98_method_expressions.singlish) | Method expressions |
| 99 | [99_closures_state.singlish](../examples/99_closures_state.singlish) | Closure state |
| 100 | [100_panic_defer.singlish](../examples/100_panic_defer.singlish) | Panic + defer interplay |
| 101 | [101_empty_interface.singlish](../examples/101_empty_interface.singlish) | Empty interface |
| 102 | [102_type_methods.singlish](../examples/102_type_methods.singlish) | Methods on named types |
| 103 | [103_string_contains.singlish](../examples/103_string_contains.singlish) | `strings.Contains` |
| 104 | [104_string_hasprefix.singlish](../examples/104_string_hasprefix.singlish) | `strings.HasPrefix` |
| 105 | [105_string_hassuffix.singlish](../examples/105_string_hassuffix.singlish) | `strings.HasSuffix` |
| 106 | [106_math_min.singlish](../examples/106_math_min.singlish) | `math.Min` |
| 107 | [107_math_max.singlish](../examples/107_math_max.singlish) | `math.Max` |
| 108 | [108_time_unix.singlish](../examples/108_time_unix.singlish) | `time.Now().Unix()` |
| 109 | [109_os_hostname.singlish](../examples/109_os_hostname.singlish) | `os.Hostname` |
| 110 | [110_bytes_reader.singlish](../examples/110_bytes_reader.singlish) | `bytes.NewReader` |
| 111 | [111_sort_float64.singlish](../examples/111_sort_float64.singlish) | `sort.Float64s` |
| 112 | [112_crypto_md5.singlish](../examples/112_crypto_md5.singlish) | `crypto/md5` |
| 113 | [113_filepath_dir.singlish](../examples/113_filepath_dir.singlish) | `filepath.Dir` |
| 114 | [114_filepath_base.singlish](../examples/114_filepath_base.singlish) | `filepath.Base` |
| 115 | [115_regexp_match.singlish](../examples/115_regexp_match.singlish) | `regexp.MatchString` |
| 116 | [116_url_escape.singlish](../examples/116_url_escape.singlish) | `url.QueryEscape` |
| 117 | [117_json_marshalindent.singlish](../examples/117_json_marshalindent.singlish) | `json.MarshalIndent` |
| 118 | [118_time_parse.singlish](../examples/118_time_parse.singlish) | `time.Parse` |
| 119 | [119_fmt_errorf.singlish](../examples/119_fmt_errorf.singlish) | `fmt.Errorf` |
| 120 | [120_sync_waitgroup_loop.singlish](../examples/120_sync_waitgroup_loop.singlish) | WaitGroup in a loop |
| 121 | [121_atomic_add.singlish](../examples/121_atomic_add.singlish) | `atomic.AddInt64` |
| 122 | [122_reflect_value.singlish](../examples/122_reflect_value.singlish) | `reflect.ValueOf` |

---

## New: Singapore Edition (123–142)

| # | File | Topics |
|---|------|--------|
| 123 | [123_string_split.singlish](../examples/123_string_split.singlish) | `strings.Split`, `strings.Join` |
| 124 | [124_multiple_assign.singlish](../examples/124_multiple_assign.singlish) | Multiple return values, variable swap |
| 125 | [125_slice_of_structs.singlish](../examples/125_slice_of_structs.singlish) | Slice of structs (kopitiam menu) |
| 126 | [126_nested_maps.singlish](../examples/126_nested_maps.singlish) | `menu[tar]menu[tar]tar` nested maps |
| 127 | [127_func_map.singlish](../examples/127_func_map.singlish) | Higher-order function, GST transform |
| 128 | [128_string_trim.singlish](../examples/128_string_trim.singlish) | `TrimSpace`, `TrimPrefix`, `TrimSuffix` |
| 129 | [129_goroutine_countdown.singlish](../examples/129_goroutine_countdown.singlish) | Goroutine countdown with WaitGroup |
| 130 | [130_map_counting.singlish](../examples/130_map_counting.singlish) | Tally map (vote counting) |
| 131 | [131_pointer_struct.singlish](../examples/131_pointer_struct.singlish) | Pointer receiver methods (bank account) |
| 132 | [132_rune_handling.singlish](../examples/132_rune_handling.singlish) | Unicode / `zhi` (rune) iteration |
| 133 | [133_type_switch.singlish](../examples/133_type_switch.singlish) | `see_how v.(type)` dispatch |
| 134 | [134_semaphore.singlish](../examples/134_semaphore.singlish) | Buffered channel semaphore |
| 135 | [135_strconv.singlish](../examples/135_strconv.singlish) | `strconv.Atoi`, `Itoa`, `ParseFloat` |
| 136 | [136_builder_pattern.singlish](../examples/136_builder_pattern.singlish) | Builder pattern / method chaining |
| 137 | [137_os_args.singlish](../examples/137_os_args.singlish) | `os.Args` with defaults |
| 138 | [138_func_dispatch.singlish](../examples/138_func_dispatch.singlish) | Map as dispatch table |
| 139 | [139_memoization.singlish](../examples/139_memoization.singlish) | Memoized Fibonacci |
| 140 | [140_interface_stack.singlish](../examples/140_interface_stack.singlish) | `interface{}` generic stack |
| 141 | [141_sort_slice.singlish](../examples/141_sort_slice.singlish) | `sort.Slice` custom comparator |
| 142 | [142_stringer_enum.singlish](../examples/142_stringer_enum.singlish) | Typed enum + `String()` (MRT lines) |

---

## Known Gotchas

> The following are **reserved Singlish keywords** and **cannot be used as variable or loop-variable names**:
> `menu`, `count`, `buat`, `upsize`, `buang`, `kwear`, `gong`, `dun`, `or`, `somemore`

> Multi-value if-init (`nasi val, ok := m[n]; ok`) is not supported. Use a separate `got val, ok = m[n]` then `nasi ok`.

> For channel receive at statement level, prefer `got _ = catch chan` over bare `catch chan` to avoid transpiler merging with the previous line.
