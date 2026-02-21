package cmd

import (
	"fmt"
	"os"
)

const transpileUsage = `Usage:
  singlish transpile <file>

Description:
  Emit the generated Go file without building.
`

func runTranspile(args []string) int {
	if len(args) == 0 || isHelpFlag(args[0]) {
		fmt.Fprint(os.Stdout, transpileUsage)
		if len(args) == 0 {
			// If run without args, usually we show usage and exit non-zero or zero?
			// US-008 says: "Negative: missing input file exits non-zero and prints a clear error"
			// But that's if I invoke it expecting a file.
			// Standard CLI: no args = usage = error?
			// But `singlish --help` prints usage and returns 0.
			// `singlish transpile` (no file) should probably error.
			if len(args) == 0 {
				fmt.Fprintln(os.Stderr, "\nError: missing input file")
				return 1
			}
		}
		return 0
	}

	inputFile := args[0]
	tempPath, err := transpileToTemp(inputFile)
	if err != nil {
		handleError(err, inputFile)
		return 1
	}

	fmt.Println(tempPath)
	return 0
}
