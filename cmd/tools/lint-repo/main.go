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

	fail := func(name string, err error) error {
		countFail++
		const format = "Linting on [FAIL](%s): %s"
		if useColor {
			ansi.Red().Printf(format, name, err).LF().Reset()
		} else {
			fmt.Printf(format, name, err)
		}
		return nil
	}

	skip := func(quiet bool, name, msg string) error {
		countSkip++
		const format = "Linting on [SKIP](%s): %s"
		if !quiet {
			if useColor {
				ansi.Yellow().Printf(format, name, msg).LF().Reset()
			} else {
				fmt.Printf(format, name, msg)
			}
		}
		return nil
	}

	pass := func(name string) error {
		countPass++
		const format = "Linting [PASS](%s)"
		if useColor {
			ansi.Green().Printf(format, name).LF().Reset()
		} else {
			fmt.Printf(format, name)
		}
		return nil
	}

	if useColor {
		ansi.Blue().
			Printf("Running Linter").
			LF().
			Reset()
	} else {
		fmt.Printf("Running Linter\n")
	}

	err := repolinter.LinterMaster(quietMode, pass, skip, fail)
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
