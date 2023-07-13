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
	"github.com/sam-caldwell/go/v2/projects/exit"
	keyvalue "github.com/sam-caldwell/go/v2/projects/keyvalue"
	"github.com/sam-caldwell/go/v2/projects/moremath"
	"github.com/sam-caldwell/go/v2/projects/repotools"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	listrepoprojects "github.com/sam-caldwell/go/v2/projects/repotools/listprojects"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
	"github.com/sam-caldwell/go/v2/projects/simpleArgs"
	"github.com/sam-caldwell/go/v2/projects/version"
	"time"
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
	var printer, color = simpleArgs.SelectPrinter()
	var banner = simpleArgs.BannerPrinter()
	var leftColumnWidth int
	var rightColumnWidth int
	var recordCount int

	simpleArgs.ExitIfHelpRequested(commandUsage)

	err = filter.FromCliArgs()
	exit.OnError(err, exit.GeneralError, commandUsage)

	projects, err = listrepoprojects.ListAsOrderedPair(filter)
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

	const displayWidthMargin = 3
	indexWidth := moremath.CountDigitsInt(recordCount)
	displayWidth := leftColumnWidth + rightColumnWidth + indexWidth + displayWidthMargin

	{
		/*
		 * Display banner if -banner is used.
		 */
		gitHash, err := repotools.GetCurrentGitHash()
		exit.OnError(err, exit.GeneralError, commandUsage)
		bannerText := fmt.Sprintf("Project Listing (version:%s git:%s)", version.Version, gitHash)
		if width := len(bannerText); width > displayWidth {
			displayWidth = width + displayWidthMargin
		}
		banner(bannerText, displayWidth, displayWidthMargin, color)
	}
	/*
	 * Walk through our project list and display each line
	 */
	for pos, project := range projects {
		name := project.Key
		manifest := project.Value.(projectmanifest.Manifest)
		printer("%*d %*s %s\n", indexWidth, pos, leftColumnWidth, manifest.GetName(), name)
	}
	{

		/*
		 * Display footer if -banner is used.
		 */
		bannerText := fmt.Sprintf("count: %d   Printed at %s", recordCount, time.Now().String())
		banner(bannerText, displayWidth, displayWidthMargin, color)
	}

	exit.OnError(err, exit.GeneralError, commandUsage)
}
