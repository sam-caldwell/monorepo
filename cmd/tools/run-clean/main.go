package main

import (
	"fmt"
	ansi2 "github.com/sam-caldwell/go/v2/projects/go/ansi"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"github.com/sam-caldwell/go/v2/projects/go/repotools"
	"github.com/sam-caldwell/go/v2/projects/go/simpleArgs"
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
		ansi2.Blue().Println("Cleaning...").Reset()
	}

	err = repotools.Clean()
	if err != nil {
		if useColor {
			ansi2.Red().Println(err.Error()).Reset()
		} else {
			fmt.Println(err)
		}
		os.Exit(exit.GeneralError)
	}
	if useColor {
		ansi2.Green().Println("ok").Reset()
	}
	os.Exit(exit.Success)
}
