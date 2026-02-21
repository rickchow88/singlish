package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	files, _ := filepath.Glob("examples/*.singlish")
	var failed []string
	for _, f := range files {
		if strings.Contains(f, "http_server") {
			continue
		}
		if strings.Contains(f, "debug_catch") {
			continue
		}
		cmd := exec.Command("./singlish", "run", f)

		// Setup timeout
		done := make(chan error, 1)
		go func() {
			done <- cmd.Run()
		}()

		select {
		case <-time.After(3 * time.Second):
			if cmd.Process != nil {
				cmd.Process.Kill()
			}
			fmt.Println("FAILED (TIMEOUT):", f)
			failed = append(failed, f)
		case err := <-done:
			if err != nil {
				fmt.Println("FAILED (ERROR):", f)
				failed = append(failed, f)
			} else {
				fmt.Println("PASSED:", f)
			}
		}
	}
	fmt.Println("\nSummary of Failures:")
	for _, f := range failed {
		fmt.Println(f)
	}
}
