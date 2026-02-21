package main

import (
	"os"

	"github.com/rickchow/singlish/cmd"
)

func main() {
	os.Exit(cmd.Execute(os.Args[1:]))
}
