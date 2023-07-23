package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	repolinter "github.com/sam-caldwell/go/v2/projects/repotools/linter"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
	"github.com/sam-caldwell/go/v2/projects/simpleArgs"
)

/*
 * run-linter
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This program creates will lint our repository
 * and its various file formats.
 *
 * See README.md
 */

import (
	"os"
)

const (
	programName  = "run-linter"
	commandUsage = `
run-linter -h|--help 
   Show usage
   
run-linter [-color]
   Run linter and if -color is present show
   the output in ANSI color.
`
	displayWidth = 40
)

// main - run-linter main function
func main() {
	var useColor = simpleArgs.UseColor()
	var quietMode = simpleArgs.QuietMode()
	var countSkip int
	var countFail int
	var countPass int

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

	banner := repocli.
		BannerMessagePrinter(programName, useColor, quietMode, displayWidth)

	notice := repocli.NoticeMessagePrinter(programName, useColor, quietMode)

	fail := repocli.FailMessagePrinter(programName, useColor, &countFail)

	skip := repocli.SkipMessagePrinter(programName, useColor, quietMode, &countSkip)

	pass := repocli.PassMessagePrinter(programName, useColor, &countPass)

	banner(ansi.Blue(), "start")

	err := repolinter.LinterMaster(notice, pass, skip, fail)
	repocli.ShowStats(programName, displayWidth, useColor, quietMode, countSkip, countFail, countSkip)
	if err != nil {
		banner(ansi.Red(), programName, "failed checks")
	}
	banner(ansi.Green(), programName, "passing all checks")
	os.Exit(exit.Success)
}
