package main

/*
 * list-projects
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This program creates will list the projects in the
 * monorepo with certain filtered characteristics.
 *
 * See README.md
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/cmd/tools/list-projects/ui"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/keyvalue/pair"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	"github.com/sam-caldwell/go/v2/projects/repotools/listprojects"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
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
	var projects pair.OrderedPair
	var printer, color = ui.SelectPrinter()
	var banner = ui.UseHeadersAndFooters()
	var sortedBy string
	var leftColumnWidth int
	var rightColumnWidth int
	var recordCount int

	exit.IfHelpRequested(commandUsage)

	err = filter.FromCliArgs()
	exit.OnError(err, exit.GeneralError, commandUsage)

	if order := ui.SortBy(); order == ui.SortByDependency {
		sortedBy = "dependency"
		projects, err = listprojects.SortedByDependencies(filter)
	} else {
		sortedBy = "name"
		projects, err = listprojects.SortedLexically(filter)
	}
	exit.OnError(err, exit.GeneralError, commandUsage)

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

	displayWidth := leftColumnWidth + rightColumnWidth + 10
	if banner {
		ui.PrintHeader(fmt.Sprintf("Project Listing (version: %s) (sort by: %s)", version.Version, sortedBy),
			color, displayWidth)
	}

	err = projects.Walk(func(key string, value interface{}) error {
		manifest := value.(projectmanifest.Manifest)
		printer(leftColumnWidth, manifest.GetName(), key)
		return nil
	})
	exit.OnError(err, exit.GeneralError, commandUsage)

	if banner {
		ui.PrintFooter(fmt.Sprintf("color: %d  date: %s", recordCount, time.Now().String()),
			color, displayWidth)
	}

	exit.OnError(err, exit.GeneralError, commandUsage)

}
