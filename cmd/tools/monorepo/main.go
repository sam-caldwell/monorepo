package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	repoBuilder "github.com/sam-caldwell/go/v2/projects/repotools/builder"
	"github.com/sam-caldwell/go/v2/projects/repotools/common"
	repoLinter "github.com/sam-caldwell/go/v2/projects/repotools/linter"
	repoScanner "github.com/sam-caldwell/go/v2/projects/repotools/scanner"
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
	helpText     = "\n\nUsage:\nbuilder [option]\n" +
		"\t --help       : display this message\n" +
		"\t --version    : show the current version\n" +
		"\t --init       : install the tools required (must be done first)\n" +
		"\t --lint       : just run the linters\n" +
		"\t --scan       : just run the security scanners\n" +
		"\t --test       : run the unit tests\n" +
		"\t --full-tests : run the integration tests (future, not working)\n" +
		"\t --build      : build all enabled projects\n\n"

	commandDirectory = "cmd"

	errAccessDenied = "Cannot access '%s': %v"
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
		var err error
		var projectName string
		if projectName, err = simpleArgs.GetOptionValue("--project"); err != nil {
			Error(err, exit.InvalidInput)
		}
		if err := repoBuilder.Build(Logf, noop, projectName); err != nil {
			Logf(ansi.Red(), "Linter failed on %s\nError:%v\n", projectName, err)
			os.Exit(exit.GeneralError)
		}
	case cmdInit:
		if err := repotools.Setup(Logf, noop); err != nil {
			Error(err, exit.GeneralError)
		}
	case cmdList:
		var filter repotools.ListProjectsFilter
		if simpleArgs.HasFlag("--disabled") {
			filter = repotools.Disabled
		}
		if simpleArgs.HasFlag("--enabled") {
			filter = repotools.Enabled
		}
		if simpleArgs.HasFlag("--all") {
			filter = repotools.All
		}

		Logf(ansi.Blue(), "Listing projects (enabled:%v)", filter.String())
		Log(ansi.Blue(), strings.Repeat("-", displayWidth))
		projectList, err := repotools.ListProjects(commandDirectory, filter)
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
		if err := repoLinter.Lint(Logf, noop, projectPath); err != nil {
			Logf(ansi.Red(), "Linter failed on %s\nError:%v\n", projectPath, err)
			os.Exit(exit.GeneralError)
		}
	case cmdScan:
		if err := repoScanner.Scan(Logf, noop, projectPath); err != nil {
			Logf(ansi.Red(), "Linter failed on %s\nError:%v\n", projectPath, err)
			os.Exit(exit.GeneralError)
		}
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
