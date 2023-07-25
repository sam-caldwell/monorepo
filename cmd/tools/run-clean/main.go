package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	"github.com/sam-caldwell/go/v2/projects/simpleArgs"
	"os"
)

const (
	programName  = "run-builds"
	displayWidth = 80
	commandUsage = `
run-clean -h|-help
	show this screen

run-clean [-color]
	Delete all build/rest artifacts
`
)

func main() {
	var useColor = simpleArgs.UseColor()
	var err error

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

	if useColor {
		ansi.Blue().Println("Cleaning...").Reset()
	}

	err = repotools.Clean()
	if err != nil {
		if useColor {
			ansi.Red().Println(err.Error()).Reset()
		} else {
			fmt.Println(err)
		}
		os.Exit(exit.GeneralError)
	}
	if useColor {
		ansi.Green().Println("ok").Reset()
	}
	os.Exit(exit.Success)
}
