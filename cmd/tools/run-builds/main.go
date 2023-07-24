package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	repoBuilder "github.com/sam-caldwell/go/v2/projects/repotools/repobuilder"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
	"github.com/sam-caldwell/go/v2/projects/simpleArgs"
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

	banner(ansi.Blue(), "Building projects")

	err = repoBuilder.Build(notice, pass, skip, fail)

	repocli.ShowStats(programName, displayWidth, useColor, quietMode, countPass, countFail, countSkip)

	if err != nil {
		banner(ansi.Red(), programName, "failed checks")
		os.Exit(exit.GeneralError)
	}
	banner(ansi.Green(), programName, "passing all checks")
	os.Exit(exit.Success)
}
