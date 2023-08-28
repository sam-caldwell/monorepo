package main

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/exit"
)

const (
	cmdBuild = "build"
	cmdInit  = "init"
	cmdLint  = "lint"
	cmdList  = "list"
	cmdScan  = "scan"

	flagVerbose = "--verbose"
	flagNoop    = "--noop"
)
const (
	displayWidth = 40
	helpText     = "\n" +
		"\nUsage:\nmonorepo {command} [options]\n"
)

func main() {
	exit.IfVersionRequested()
	//var command string
	//var verbose bool
	//var noop bool
	//var filter filters.Filter
	//
	//command = simpleArgs.GetCommand(helpText)
	//verbose = simpleArgs.HasFlag(flagVerbose)
	//noop = simpleArgs.HasFlag(flagNoop)
	//
	//Log, Logf, Error := simpleLogger.Logger(!verbose)
	//
	//Log(ansi.Blue(), "Starting")
	//
	//if command != cmdInit {
	//	if err := filter.FromCliArgs(); err != nil {
	//		Error(err, exit.GeneralError)
	//	}
	//}
	//
	//switch command {
	//case cmdBuild:
	//	var projectName string
	//	if err := repoBuilder.Build(Logf, noop); err != nil {
	//		Logf(ansi.Red(), "Linter failed on %s\nError:%v\n", projectName, err)
	//		os.Exit(exit.GeneralError)
	//	}
	//case cmdInit:
	//	if err := repotools.Setup(Logf, noop); err != nil {
	//		Error(err, exit.GeneralError)
	//	}
	//case cmdList:
	//	Logf(ansi.Blue(), "Listing projects (flags:%s)\n", strings.Join(filter.String(), ", "))
	//	Log(ansi.Blue(), strings.Repeat("-", displayWidth))
	//	projectList, err := listrepoprojects.ListProjects(filter)
	//	if err != nil {
	//		Error(err, exit.GeneralError)
	//	}
	//	count := 0
	//	for _, project := range projectList {
	//		fmt.Println(project)
	//		count++
	//	}
	//	Log(ansi.Green(), strings.Repeat("-", displayWidth))
	//	Logf(ansi.Green(), "count: %d\n", count)
	//	Log(ansi.Green(), strings.Repeat("-", displayWidth))
	//	return
	//case cmdLint:
	//	//if err := repoLinter.Lint(Logf, noop, projectPath); err != nil {
	//	//	Logf(ansi.Red(), "Linter failed on %s\nError:%v\n", projectPath, err)
	//	//	os.Exit(exit.GeneralError)
	//	//}
	//case cmdScan:
	//	//if err := repoScanner.Scan(Logf, noop, projectPath); err != nil {
	//	//	Logf(ansi.Red(), "Linter failed on %s\nError:%v\n", projectPath, err)
	//	//	os.Exit(exit.GeneralError)
	//	//}
	//default:
	//	Error("Unknown command: "+command, exit.InvalidInput)
	//}

	/*
	 * build all enabled projects
	 */
	//if *build {
	//	return
	//}
}
