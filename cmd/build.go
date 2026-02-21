package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const buildUsage = `Usage:
  singlish build <file>

Description:
  Transpile and build a binary from a .singlish file.
`

func runBuild(args []string) int {
	if len(args) == 0 || isHelpFlag(args[0]) {
		fmt.Fprint(os.Stdout, buildUsage)
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "\nError: missing input file")
			return 1
		}
		return 0
	}

	inputFile := args[0]
	tempPath, err := transpileToTemp(inputFile)
	if err != nil {
		handleError(err, inputFile)
		return 1
	}
	defer os.Remove(tempPath)

	// Determine output name
	base := filepath.Base(inputFile)
	ext := filepath.Ext(base)
	outputName := strings.TrimSuffix(base, ext)
	if outputName == "" {
		outputName = "main"
	}

	// Run go build
	cmd := exec.Command("go", "build", "-o", outputName, tempPath)
	cmd.Stdout = os.Stdout
	// Stderr handled by executeWithInsults

	if err := executeWithInsults(cmd); err != nil {
		return 1
	}

	return 0
}
