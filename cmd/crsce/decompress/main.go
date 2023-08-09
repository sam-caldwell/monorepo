package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/crsce"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/simpleArgs"
	"os"
)

const (
	programName  = "compress"
	commandUsage = `
decompress -h|-help
	show this screen

decompress sourceFileName destinationFileName
    decompress sourceFileName and write to destinationFileName
`
)

func main() {
	var err error
	var c crsce.Crsce
	var sourceFile string
	var destinationFile string

	exit.OnCondition(len(os.Args) < 3, exit.MissingArg, "Missing input(s)", commandUsage)
	exit.IfHelpRequested(commandUsage)

	sourceFile, err = simpleArgs.GetOptionValue("--in")
	exit.OnError(err, exit.MissingArg, commandUsage)

	destinationFile, err = simpleArgs.GetOptionValue("--out")
	exit.OnError(err, exit.MissingArg, commandUsage)

	if err = c.Open(sourceFile); err != nil {
		fmt.Printf("Cannot open source file.  Error: %v", err)
		os.Exit(exit.GeneralError)
	}
	if err = c.Decompress(destinationFile); err != nil {
		fmt.Printf("Cannot decompress to destination file.  Error: %v", err)
		os.Exit(exit.GeneralError)
	}
}
