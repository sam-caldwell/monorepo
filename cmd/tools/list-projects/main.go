package main

/*
 * list-projects
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This program creates will list the projects in the
 * monorepo with certain filtered characteristics.
 *
 * Warning: Because the project list is a key-value
 * store and because keyvalue is based on a map order
 * is **NOT** guaranteed.
 *
 * See README.md
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	keyvalue "github.com/sam-caldwell/go/v2/projects/keyvalue"
	"github.com/sam-caldwell/go/v2/projects/moremath"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	listrepoprojects "github.com/sam-caldwell/go/v2/projects/repotools/listprojects"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
	"github.com/sam-caldwell/go/v2/projects/simpleArgs"
	"github.com/sam-caldwell/go/v2/projects/version"
	"os"
)

const (
	commandUsage = `
list-projects [display flags] [filter flags]

    display flags:
        -color   // Show the output in color
        -banner  // Pretty print the output with a banner. 

    filter flags:
        -commands      // Show only command projects
        -buildEnabled  // Show only build-enabled projects
        -lintEnabled  // Show only lint-enabled projects
        -scanEnabled  // Show only scan-enabled projects (e.g. security scanners)
        -signEnabled  // Show only sign-enabled projects (e.g. projects ready to code-sign)
        -packEnabled  // Show only packaging-enabled projects (e.g. projects ready for packaging)
        -os <opsys>  // Show only projects which support the given operating system
        -arch <arch>  // Show only projects which support the given CPU Architecture
`
)

// main - list-project main function
func main() {
	var err error
	var filter filters.Filter
	var projects []keyvalue.OrderedPair
	var printer, color = selectPrinter()
	var banner = useHeadersAndFooters()
	var leftColumnWidth int
	var rightColumnWidth int
	var recordCount int

	simpleArgs.ExitIfHelpRequested(commandUsage)

	err = filter.FromCliArgs()
	exit.OnError(err, exit.GeneralError, commandUsage)

	projects, err = func() ([]keyvalue.OrderedPair, error) {
		raw, e := listrepoprojects.ListProjects(filter)
		return raw.ToOrderedList(true), e
	}()
	exit.OnError(err, exit.GeneralError, commandUsage)

	/*
	 * count our records (for the optional banner)
	 * and calculate our column widths (left and right)
	 */
	for _, project := range projects {
		name := project.Key
		manifest := project.Value.(projectmanifest.Manifest)
		if width := len(manifest.GetName()); width > leftColumnWidth {
			leftColumnWidth = width
		}
		if width := len(name); width > rightColumnWidth {
			rightColumnWidth = width
		}
		recordCount++
	}

	/*
	 * Display banner if -banner is used.
	 */
	const displayWidthMargin = 3
	indexWidth := moremath.CountDigitsInt(recordCount)
	displayWidth := leftColumnWidth + rightColumnWidth + indexWidth + displayWidthMargin
	if banner {
		gitHash, err := repotools.GetCurrentGitHash()
		exit.OnError(err, exit.GeneralError, commandUsage)
		bannerText := fmt.Sprintf("Project Listing (version:%s git:%s)", version.Version, gitHash)
		if width := len(bannerText); width > displayWidth {
			displayWidth = width + displayWidthMargin
		}
		if color {
			ansi.Blue().
				Line("-", displayWidth).
				Space().Print(bannerText).LF().
				Line("-", displayWidth).
				Reset()
		} else {
			fmt.Println(bannerText)
			fmt.Println("------------------")
		}
	}
	/*
	 * Walk through our project list and display each line
	 */
	for pos, project := range projects {
		name := project.Key
		manifest := project.Value.(projectmanifest.Manifest)
		printer(indexWidth, pos, leftColumnWidth, manifest.GetName(), name)
	}
	/*
	 * Display footer if -banner is used.
	 */
	if banner {
		if color {
			ansi.Blue().
				LF().
				Line("-", displayWidth).
				Space().Printf("count: %d", recordCount).LF().
				Line("-", displayWidth).
				Reset()
		} else {
			fmt.Println("------------------")
			fmt.Printf("count: %d\n", recordCount)

		}
	}

	exit.OnError(err, exit.GeneralError, commandUsage)

}

// printerFunction - a simple function pattern for our Printers
type printerFunction func(indexWidth, pos, width int, name, key string)

// SelectPrinter - select a color or non-color printer
func selectPrinter() (printerFunction, bool) {
	for _, arg := range os.Args {
		if arg == "-color" {
			return func(indexWidth, pos, width int, name, key string) {
				ansi.Blue().
					Space().
					Printf("%0*d %*s", indexWidth, pos, width, name).
					Space().
					Yellow().
					Printf("%s", key).LF().
					Reset()
			}, true
		}
	}
	return func(indexWidth, pos, width int, name, key string) {
		fmt.Printf(" %0*d %*s %s\n", indexWidth, pos, width, name, key)
	}, false
}

// UseHeadersAndFooters determine if we should pretty print
func useHeadersAndFooters() bool {
	for _, arg := range os.Args {
		if arg == "-banner" {
			return true
		}
	}
	return false
}
