package main

/*
 * list-projects
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This program creates will list the projects in the
 * monorepo with certain filtered characteristics.
 *
 * Warning: Because the project list is a key-value
 * store and because KeyValue is based on a map order
 * is **NOT** guaranteed.
 *
 * See README.md
 */

import (
	"fmt"
	keyvalue "github.com/sam-caldwell/go/v2/projects/KeyValue"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	listrepoprojects "github.com/sam-caldwell/go/v2/projects/repotools/listrepoprojects"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
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
	var projects keyvalue.KeyValue
	var printer, color = selectPrinter()
	var banner = useHeadersAndFooters()
	var leftColumnWidth int
	var rightColumnWidth int
	var recordCount int

	exitIfHelpRequested()

	err = filter.FromCliArgs()
	exit.OnError(err, exit.GeneralError, commandUsage)

	projects, err = listrepoprojects.ListProjects(filter)
	exit.OnError(err, exit.GeneralError, commandUsage)

	/*
	 * count our records (for the optional banner)
	 * and calculate our column widths (left and right)
	 */
	err = projects.Walk(func(key string, value interface{}) error {
		manifest := value.(projectmanifest.Manifest)
		if width := len(manifest.GetName()); width > leftColumnWidth {
			leftColumnWidth = width
		}
		if width := len(key); width > rightColumnWidth {
			rightColumnWidth = width
		}
		recordCount++
		return nil
	})
	exit.OnError(err, exit.GeneralError, commandUsage)
	/*
	 * Display banner if -banner is used.
	 */
	const displayWidthMargin = 2
	displayWidth := leftColumnWidth + rightColumnWidth + displayWidthMargin
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
	err = projects.Walk(func(key string, value interface{}) error {
		manifest := value.(projectmanifest.Manifest)
		printer(leftColumnWidth, manifest.GetName(), key)
		return nil
	})
	exit.OnError(err, exit.GeneralError, commandUsage)
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

// exitIfHelpRequested - intercept a request for help information
func exitIfHelpRequested() {
	for _, arg := range os.Args {
		exit.OnCondition(arg == "-h" || arg == "--help", 0, "", commandUsage)
	}
}

// printerFunction - a simple function pattern for our Printers
type printerFunction func(width int, name, key string)

// SelectPrinter - select a color or non-color printer
func selectPrinter() (printerFunction, bool) {
	for _, arg := range os.Args {
		if arg == "-color" {
			return func(width int, name, key string) {
				ansi.Blue().
					Printf("%*s", width, name).
					Space().
					Yellow().
					Printf("%s", key).LF().
					Reset()
			}, true
		}
	}
	return func(width int, name, key string) {
		fmt.Printf("%*s %s\n", width, name, key)
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
