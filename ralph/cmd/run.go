package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

const runUsage = `Usage:
  singlish run <file>

Description:
  Transpile and run a .singlish file.
`

func runRun(args []string) int {
	if len(args) == 0 || isHelpFlag(args[0]) {
		fmt.Fprint(os.Stdout, runUsage)
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "\nError: missing input file")
			return 1
		}
		return 0
	}

	inputFile := args[0]
	tempPath, err := transpileToTemp(inputFile)
	if err != nil {
		printErrorWithInsult(err)
		return 1
	}
	defer os.Remove(tempPath)

	// Run go run
	// Pass remaining args if any (though US-008 doesn't strictly require it, it's good practice)
	// But go run syntax is `go run [build flags] <files> [arguments...]`
	goArgs := append([]string{"run", tempPath}, args[1:]...)
	cmd := exec.Command("go", goArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	// Stderr handled by executeWithInsults

	if err := executeWithInsults(cmd); err != nil {
		return 1
	}

	return 0
}
