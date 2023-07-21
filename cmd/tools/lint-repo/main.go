package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
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
)

// main - lint-repo main function
func main() {
	var useColor = simpleArgs.UseColor()

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

	fail := func(name string, err error) {
		const format = "Linting on [FAIL](%s): %s"
		if useColor {
			ansi.Blue().Printf(format, name, err).LF().Reset()
		} else {
			fmt.Printf(format, name, err)
		}
	}

	pass := func(name string) {
		const format = "Linting [PASS](%s)"
		if useColor {
			ansi.Blue().Printf(format, name).LF().Reset()
		} else {
			fmt.Printf(format, name)
		}
	}

	if useColor {
		ansi.Blue().Printf("Running Linter").LF().Reset()
	} else {
		fmt.Printf("Running Linter\n")
	}

	if err := repolinter.LinterMaster(pass, fail); err != nil {
		if useColor {
			ansi.Red().Printf("Linter failed (%s)", err).LF().Reset()
			os.Exit(exit.Success)
		} else {
			fmt.Printf("Linter failed (%s)\n", err)
			os.Exit(exit.GeneralError)
		}
	}
}
