package main

import (
	ansi "github.com/sam-caldwell/go/v2/projects/go/ansi"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"github.com/sam-caldwell/go/v2/projects/go/repotools/repotester"
	repocli2 "github.com/sam-caldwell/go/v2/projects/go/repotools/ui"
	simpleArgs2 "github.com/sam-caldwell/go/v2/projects/go/simpleArgs"
	"os"
)

const (
	programName  = "run-tests"
	displayWidth = 80
	commandUsage = `
run-tests -h|-help
	show this screen

run-tests [-color]
	Run the tests for every project in the repo, then run tests
    for every command defined in the repo.
`
)

func main() {
	var useColor = simpleArgs2.UseColor()
	var quietMode = simpleArgs2.QuietMode()
	var countSkip int
	var countFail int
	var countPass int
	var err error

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

	banner := repocli2.BannerMessagePrinter(programName, useColor, quietMode, displayWidth)

	notice := repocli2.NoticeMessagePrinter(programName, useColor, quietMode)

	fail := repocli2.FailMessagePrinter(programName, useColor, &countFail)

	skip := repocli2.SkipMessagePrinter(programName, useColor, quietMode, &countSkip)

	pass := repocli2.PassMessagePrinter(programName, useColor, &countPass)

	testRunner := repotester.Setup(notice, pass, skip, fail)

	for _, testGroup := range []string{"projects", "cmd"} {
		banner(ansi.Blue(), programName+": start "+testGroup)
		err = testRunner(testGroup)
		if err != nil {
			fail("testType:"+testGroup, "outcome:failure", err)
			break
		}
	}
	repocli2.ShowStats(programName, displayWidth, useColor, quietMode, countPass, countFail, countSkip)
	if err != nil {
		banner(ansi.Red(), programName, "failed checks")
		os.Exit(exit.GeneralError)
	}
	banner(ansi.Green(), programName, "passing all checks")
	os.Exit(exit.Success)
}
