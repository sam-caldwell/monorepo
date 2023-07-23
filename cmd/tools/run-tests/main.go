package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
	"github.com/sam-caldwell/go/v2/projects/simpleArgs"
)

const (
	programName  = "run-linter"
	displayWidth = 80
	commandUsage = `
run-tests -h|-help
	show this screen

run-tests [-color] -project <project_name>
	Run the tests for a given project and its dependencies.

run-tests [-color]
	Run the tests for every project in the repo, then run tests
    for every command defined in the repo.
`
)

func main() {
	var useColor = simpleArgs.UseColor()
	var quietMode = simpleArgs.QuietMode()
	//var countSkip int
	//var countFail int
	//var countPass int

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

	banner := repocli.
		BannerMessagePrinter(programName, useColor, quietMode, displayWidth)
	//notice := repocli.
	//	NoticeMessagePrinter(programName, useColor, quietMode)
	//
	//fail := repocli.FailMessagePrinter(programName, useColor, quietMode, &countFail)
	//
	//skip := repocli.SkipMessagePrinter(programName, useColor, quietMode, &countSkip)
	//
	//pass := repocli.PassMessagePrinter(programName, useColor, quietMode, &countPass)

	banner(ansi.Blue(), "start")
}
