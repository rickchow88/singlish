package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	files := map[string]string{
		"examples/01_hello_world.singlish":            "// Hello World in Singlish",
		"examples/02_fizzbuzz.singlish":               "// FizzBuzz Implementation",
		"examples/03_coffee_shop_queue.singlish":      "// Coffee Shop Queue Simulation",
		"examples/04_structs_and_interfaces.singlish": "// Structs and Interfaces",
		"examples/05_error_handling.singlish":         "// Error Handling",
		"examples/08_defer.singlish":                  "// Defer execution order",
		"examples/09_switch_case.singlish":            "// Switch Statements",
		"examples/100_panic_defer.singlish":           "// Panic and Recovery with Defer",
		"examples/101_empty_interface.singlish":       "// Empty Interfaces",
		"examples/102_type_methods.singlish":          "// Type Methods",
		"examples/10_panic_recover.singlish":          "// Panic and Recover",
		"examples/12_constants.singlish":              "// Constants",
		"examples/26_mutex.singlish":                  "// Mutex Locks",
		"examples/28_atomic.singlish":                 "// Atomic Operations",
		"examples/31_lucky_numbers.singlish":          "// Lucky Numbers",
		"examples/32_makan_order.singlish":            "// Makan Order Processing",
		"examples/93_sync_cond.singlish":              "// sync.Cond usage",
		"examples/94_channel_range.singlish":          "// Range over channels",
		"examples/96_embedded_interfaces.singlish":    "// Embedded Interfaces",
		"examples/97_struct_tags.singlish":            "// Struct Tags",
		"examples/98_method_expressions.singlish":     "// Method Expressions",
		"examples/99_closures_state.singlish":         "// Closures and State",
	}

	for path, comment := range files {
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading", path, err)
			continue
		}

		if !bytes.HasPrefix(content, []byte("//")) {
			newContent := append([]byte(comment+"\n"), content...)
			err = os.WriteFile(path, newContent, 0644)
			if err != nil {
				fmt.Println("Error writing", path, err)
			} else {
				fmt.Println("Added comment to", path)
			}
		}
	}
}
