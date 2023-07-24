package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	repotester "github.com/sam-caldwell/go/v2/projects/repotools/tester"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
	"github.com/sam-caldwell/go/v2/projects/simpleArgs"
	"os"
)

const (
	programName  = "run-tests"
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
	var countSkip int
	var countFail int
	var countPass int
	var err error

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

	banner := repocli.
		BannerMessagePrinter(programName, useColor, quietMode, displayWidth)

	notice := repocli.NoticeMessagePrinter(programName, useColor, quietMode)

	fail := repocli.FailMessagePrinter(programName, useColor, &countFail)

	skip := repocli.SkipMessagePrinter(programName, useColor, quietMode, &countSkip)

	pass := repocli.PassMessagePrinter(programName, useColor, &countPass)

	testRunner := repotester.Setup(notice, pass, skip, fail)

	for _, testGroup := range []string{"projects", "cmd"} {
		banner(ansi.Blue(), programName+": start "+testGroup)
		err = testRunner(testGroup)
		if err != nil {
			fail("testType:"+testGroup, "outcome:failure", err)
			break
		}
	}
	repocli.ShowStats(programName, displayWidth, useColor, quietMode, countPass, countFail, countSkip)
	if err != nil {
		banner(ansi.Red(), programName, "failed checks")
	}
	banner(ansi.Green(), programName, "passing all checks")
	os.Exit(exit.Success)
}
