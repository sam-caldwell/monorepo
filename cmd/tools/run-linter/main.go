package main

import (
	ansi2 "github.com/sam-caldwell/go/v2/projects/go/ansi"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"github.com/sam-caldwell/go/v2/projects/go/repotools/linter"
	repocli2 "github.com/sam-caldwell/go/v2/projects/go/repotools/ui"
	simpleArgs2 "github.com/sam-caldwell/go/v2/projects/go/simpleArgs"
	"os"
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
	var useColor = simpleArgs2.UseColor()
	var quietMode = simpleArgs2.QuietMode()
	var countSkip int
	var countFail int
	var countPass int

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

	banner := repocli2.BannerMessagePrinter(programName, useColor, quietMode, displayWidth)

	notice := repocli2.NoticeMessagePrinter(programName, useColor, quietMode)

	fail := repocli2.FailMessagePrinter(programName, useColor, &countFail)

	skip := repocli2.SkipMessagePrinter(programName, useColor, quietMode, &countSkip)

	pass := repocli2.PassMessagePrinter(programName, useColor, &countPass)

	banner(ansi2.Blue(), "start")

	err := repolinter.LinterMaster(useColor, notice, pass, skip, fail)
	repocli2.ShowStats(programName, displayWidth, useColor, quietMode, countPass, countFail, countSkip)
	if err != nil {
		banner(ansi2.Red(), programName, "failed checks")
	}
	banner(ansi2.Green(), programName, "passing all checks")
	os.Exit(exit.Success)
}
