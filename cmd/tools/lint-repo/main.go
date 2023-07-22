package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/misc"
	repolinter "github.com/sam-caldwell/go/v2/projects/repotools/linter"
	"github.com/sam-caldwell/go/v2/projects/simpleArgs"
)

/*
 * lint-repo
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This program creates will lint our repository
 * and its various file formats.
 *
 * See README.md
 */

import (
	"fmt"
	"os"
)

const (
	commandUsage = `
lint-repo -h|--help 
   Show usage
   
lint-repo [-color]
   Run linter and if -color is present show
   the output in ANSI color.
`
	displayWidth = 40
)

// main - lint-repo main function
func main() {
	var useColor = simpleArgs.UseColor()
	var quietMode = simpleArgs.QuietMode()
	var countSkip int
	var countFail int
	var countPass int

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

	notice := func(format string, args ...any) {
		if quietMode {
			return
		}
		if useColor {
			ansi.Yellow().Printf(">>"+format, args...).LF().Reset()
		} else {
			fmt.Printf(format+"\n", args...)
		}
	}

	fail := func(name, linter string, err error) {
		countFail++
		const format = "Linting (%s) [FAIL](%s): %v"
		if useColor {
			ansi.Red().Printf(format, linter, name, err).LF().Reset()
		} else {
			fmt.Printf(format, linter, name, err)
		}
	}

	skip := func(name, linter, msg string) {
		countSkip++
		const format = "Linting (%s) [SKIP](%s): %s"
		if !quietMode {
			if useColor {
				ansi.Yellow().Printf(format, linter, name, msg).LF().Reset()
			} else {
				fmt.Printf(format, linter, name, msg)
			}
		}
	}

	pass := func(name, linter string) {
		countPass++
		const format = "Linting (%s) [PASS](%s)"
		if useColor {
			ansi.Green().Printf(format, linter, name).LF().Reset()
		} else {
			fmt.Printf(format, linter, name)
		}
	}

	if useColor {
		ansi.Blue().
			Printf("Running Linter").
			LF().
			Reset()
	} else {
		fmt.Printf("Running Linter\n")
	}

	err := repolinter.LinterMaster(notice, pass, skip, fail)
	misc.ShowStats(useColor,
		"Linter Stats",
		fmt.Sprintf("  Total:%6d", countPass+countSkip+countFail),
		map[string]int{
			"pass": countPass,
			"fail": countFail,
			"skip": countSkip,
		})
	if err != nil {
		if useColor {
			ansi.
				Red().
				Line("-", displayWidth).
				Printf("Linter failed (%s)", err).
				LF().
				Line("-", displayWidth).
				Reset()
			os.Exit(exit.Success)
		} else {
			fmt.Printf("\nLinter failed (%s)\n", err)
			os.Exit(exit.GeneralError)
		}
	}
	if useColor {
		ansi.Green().
			Line("-", displayWidth).
			Printf("Linter passing").
			LF().
			Line("-", displayWidth).
			Reset()
		os.Exit(exit.Success)
	} else {
		fmt.Printf("\nLinter passing\n")
		os.Exit(exit.GeneralError)
	}
}
