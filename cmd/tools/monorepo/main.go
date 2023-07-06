package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	repoBuilder "github.com/sam-caldwell/go/v2/projects/repotools/builder"
	"github.com/sam-caldwell/go/v2/projects/repotools/common"
	"github.com/sam-caldwell/go/v2/projects/repotools/common/projectFilter"
	"github.com/sam-caldwell/go/v2/projects/simpleArgs"
	"github.com/sam-caldwell/go/v2/projects/simpleLogger"
	"os"
	"strings"
)

const (
	cmdBuild = "build"
	cmdInit  = "init"
	cmdLint  = "lint"
	cmdList  = "list"
	cmdScan  = "scan"
)
const (
	displayWidth = 40
	helpText     = "\n" +
		"\nUsage:\nmonorepo {command} [options]\n"
)

func main() {
	var command string
	var verbose bool
	var noop bool

	command = simpleArgs.GetCommand(helpText)
	verbose = simpleArgs.HasFlag("--verbose")
	noop = simpleArgs.HasFlag("--noop")

	Log, Logf, Error := simpleLogger.Logger(!verbose)
	Log(ansi.Blue(), "Starting")

	switch command {
	case cmdBuild:
		var projectName string
		if err := repoBuilder.Build(Logf, noop); err != nil {
			Logf(ansi.Red(), "Linter failed on %s\nError:%v\n", projectName, err)
			os.Exit(exit.GeneralError)
		}
	case cmdInit:
		if err := repotools.Setup(Logf, noop); err != nil {
			Error(err, exit.GeneralError)
		}
	case cmdList:
		var filter projectFilter.Filter
		if err := filter.FromCliArgs(); err != nil {
			Error(err, exit.GeneralError)
		}

		Logf(ansi.Blue(), "Listing projects (enabled:%v)", filter.String())
		Log(ansi.Blue(), strings.Repeat("-", displayWidth))
		projectList, err := repotools.ListProjects(filter)
		if err != nil {
			Error(err, exit.GeneralError)
		}
		count := 0
		for _, project := range projectList {
			fmt.Println(project)
			count++
		}
		Log(ansi.Green(), strings.Repeat("-", displayWidth))
		Logf(ansi.Green(), "count: %d", count)
		Log(ansi.Green(), strings.Repeat("-", displayWidth))
		return
	case cmdLint:
		//if err := repoLinter.Lint(Logf, noop, projectPath); err != nil {
		//	Logf(ansi.Red(), "Linter failed on %s\nError:%v\n", projectPath, err)
		//	os.Exit(exit.GeneralError)
		//}
	case cmdScan:
		//if err := repoScanner.Scan(Logf, noop, projectPath); err != nil {
		//	Logf(ansi.Red(), "Linter failed on %s\nError:%v\n", projectPath, err)
		//	os.Exit(exit.GeneralError)
		//}
	default:
		Error("Unknown command: "+command, exit.InvalidInput)
	}

	/*
	 * build all enabled projects
	 */
	//if *build {
	//	return
	//}
}
