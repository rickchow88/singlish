package cmd

import (
	"fmt"
	"os"
	"strings"
)

var DictionaryPath string

const usageBanner = `singlish - Singlish to Go CLI

Usage:
  singlish [global flags] <command> [args]

Global Flags:
  --dictionary <path>   Path to dictionary file (default: dictionary.txt)

Commands:
  build       Transpile and build a binary from a .singlish file
  fmt         Format Singlish source code
  run         Transpile and run a .singlish file
  transpile   Emit the generated Go file without building

Use "singlish <command> --help" for more information about a command.
`

func Execute(args []string) int {
	// Parse global flags
	var cleanArgs []string
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--dictionary" {
			if i+1 < len(args) {
				DictionaryPath = args[i+1]
				i++ // skip value
			} else {
				fmt.Fprintln(os.Stderr, "Error: --dictionary flag requires an argument")
				return 1
			}
		} else if strings.HasPrefix(arg, "--dictionary=") {
			DictionaryPath = strings.TrimPrefix(arg, "--dictionary=")
		} else {
			cleanArgs = append(cleanArgs, arg)
		}
	}
	args = cleanArgs

	if len(args) == 0 || isHelpFlag(args[0]) {
		printUsage(os.Stdout)
		return 0
	}

	switch args[0] {
	case "build":
		return runBuild(args[1:])
	case "fmt":
		return runFmt(args[1:])
	case "run":
		return runRun(args[1:])
	case "transpile":
		return runTranspile(args[1:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", args[0])
		return 1
	}
}

func printUsage(out *os.File) {
	fmt.Fprint(out, usageBanner)
}

func isHelpFlag(arg string) bool {
	return arg == "-h" || arg == "--help" || strings.EqualFold(arg, "help")
}
