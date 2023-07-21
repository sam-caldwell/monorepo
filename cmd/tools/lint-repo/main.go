package main

import "github.com/sam-caldwell/go/v2/projects/exit"

/*
 * lint-repo
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This program creates will lint our repository
 * and its various file formats.
 *
 * See README.md
 */
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

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

}
