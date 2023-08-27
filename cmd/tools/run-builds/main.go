package main

import (
	ansi "github.com/sam-caldwell/go/v2/projects/go/ansi"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	repoBuilder "github.com/sam-caldwell/go/v2/projects/go/repotools/repobuilder"
	repocli2 "github.com/sam-caldwell/go/v2/projects/go/repotools/ui"
	simpleArgs2 "github.com/sam-caldwell/go/v2/projects/go/simpleArgs"
	"os"
)

const (
	programName  = "run-builds"
	displayWidth = 80
	commandUsage = `
run-builds -h|-help
	show this screen

run-builds [-color]
	Run the builds for every command project in the repo.
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

	banner(ansi.Blue(), "Building projects")

	err = repoBuilder.Build(notice, pass, skip, fail)

	repocli2.ShowStats(programName, displayWidth, useColor, quietMode, countPass, countFail, countSkip)

	if err != nil {
		banner(ansi.Red(), programName, "failed checks")
		os.Exit(exit.GeneralError)
	}
	banner(ansi.Green(), programName, "passing all checks")
	os.Exit(exit.Success)
}
