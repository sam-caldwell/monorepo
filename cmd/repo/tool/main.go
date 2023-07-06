package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	repoBuilder "github.com/sam-caldwell/go/v2/projects/repotools/builder"
	"github.com/sam-caldwell/go/v2/projects/repotools/common"
	repoLinter "github.com/sam-caldwell/go/v2/projects/repotools/linter"
	repoScanner "github.com/sam-caldwell/go/v2/projects/repotools/scanner"
	repoTester "github.com/sam-caldwell/go/v2/projects/repotools/tester"
	"github.com/sam-caldwell/go/v2/projects/simpleLogger"
	programVersion "github.com/sam-caldwell/go/v2/projects/version"
	"os"
	"strings"
)

const (
	cmdUsage = "\n\nUsage:\nbuilder [option]\n" +
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
	/*
	 * Command-line arguments
	 */
	build := flag.Bool("build", false, "build project(s)")
	disabled := flag.Bool("disabled", false, "disabled projects only")
	enabled := flag.Bool("enabled", false, "enabled projects only")
	init := flag.Bool("init", false, "initialize this tool and install all dependencies")
	lint := flag.Bool("lint", false, "lint project(s)")
	list := flag.Bool("list", false, "list projects")
	noop := flag.Bool("noop", false, "do nothing, but print messages about what you would have done.")
	project := flag.String("project", words.EmptyString, "operate on a specific project")
	scan := flag.Bool("scan", false, "scan project(s)")
	test := flag.Bool("test", false, "test project(s)")
	verbose := flag.Bool("verbose", false, "print verbose messaging")
	version := flag.Bool("version", false, "print version")
	flag.Parse()
	/*
	 * Set up a logger function
	 */
	Log, Logf, Error := simpleLogger.Logger(!*verbose)
	/*
	 * wrapped functions because only sloppy people use globals.
	 */
	buildNotExpected := func() {
		if *build {
			Error("-build not expected here", exit.InvalidInput)
		}
	}
	enabledDisabledNotExpected := func() {
		if *enabled || *disabled {
			Error("-enabled and -disabled not expected here", exit.InvalidInput)
		}
	}
	initNotExpected := func() {
		if *init {
			Error("-init not expected here", exit.InvalidInput)
		}
	}
	lintNotExpected := func() {
		if *lint {
			Error("-lint not expected here", exit.InvalidInput)
		}
		enabledDisabledNotExpected()
	}
	listNotExpected := func() {
		if *list {
			Error("-list not expected here", exit.InvalidInput)
		}
	}
	projectNotExpected := func() {
		if *project != words.EmptyString {
			Error("-project not expected here", exit.InvalidInput)
		}
	}
	scanNotExpected := func() {
		if *scan {
			Error("-scan not expected here", exit.InvalidInput)
		}
	}
	testNotExpected := func() {
		if *test {
			Error("-test not expected here", exit.InvalidInput)
		}
	}
	versionNotExpected := func() {
		if *version {
			Error("-version not expected here", exit.InvalidInput)
		}
	}
	noopNotExpected := func() {
		if *noop {
			Error("-noop not expected here", exit.InvalidInput)
		}
	}
	/*
	 * start of the real main function
	 */
	Log(ansi.Blue(), "Starting")
	/*
	 * handle list / version / *init together since they have similar
	 * expectations
	 */
	if *version || *list || *init {
		buildNotExpected()
		initNotExpected()
		lintNotExpected()
		projectNotExpected()
		scanNotExpected()
		testNotExpected()
		/*
		 * initialize the program's dependencies.
		 */
		if *init {
			Log(ansi.Blue(), "Initializing the tool")
			if err := repotools.Setup(*noop); err != nil {
				Error(err, exit.GeneralError)
			}

			if err := repoLinter.Setup(*noop); err != nil {
				Error(err, exit.GeneralError)
			}

			if err := repoScanner.Setup(*noop); err != nil {
				Error(err, exit.GeneralError)
			}

			if err := repoBuilder.Setup(*noop); err != nil {
				Error(err, exit.GeneralError)
			}

			if err := repoTester.Setup(*noop); err != nil {
				Error(err, exit.GeneralError)
			}
			Log(ansi.Green(), "Initializing ok")
			return
		}
		noopNotExpected()
		/*
		 * show program version
		 */
		if *version {
			ansi.Println(programVersion.Version).Fatal(0)
		}
		versionNotExpected()
		/*
		 * generate a project list
		 */
		if *list {
			buildNotExpected()

			listProjects := func(enabled bool) {
				const displayWidth = 40
				Logf(ansi.Blue(), "Listing projects (enabled:%v)", enabled)
				Log(ansi.Blue(), strings.Repeat("-", displayWidth))
				projectList, err := repotools.GetProjects(commandDirectory, enabled)
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
			} /* end of function */

			if *enabled && *disabled {
				Error("Error: specify only -enabled or -disabled but not both", exit.GeneralError)
				return
			}
			if *enabled {
				listProjects(true)
				return
			}
			if *disabled {
				listProjects(false)
				return
			}
			Error("Error: specify -enabled or -disabled", exit.GeneralError)
		}
		listNotExpected()
	}
	listNotExpected()
	versionNotExpected()
	initNotExpected()
	/*
	 * Get our project listing for any further operations one time
	 * so we don't reinvent wheels.
	 */
	var err error
	var projectList []string

	if *project == "" {
		projectList, err = repotools.GetProjects("./", true)
		if err != nil {
			Error(err, exit.GeneralError)
		}
	} else {
		projectList = []string{*project}
	}
	/*
	 * lint all enabled projects
	 */
	if *lint {
		for _, projectPath := range projectList {
			Logf(ansi.Yellow(), "Linting %s...", projectPath)
			if err := repoLinter.Lint(*noop, projectPath); err != nil {
				Logf(ansi.Red(), "Linter failed on %s\nError:%v\n", projectPath, err)
				os.Exit(exit.GeneralError)
			}
			Logf(ansi.Green(), "Lint ok: %s", projectPath)
		}
		Log(ansi.Green(), "Linting successful")
		return
	}
	/*
	 * scan all enabled projects
	 */
	if *scan {
		for _, projectPath := range projectList {
			Logf(ansi.Yellow(), "Scanning %s...", projectPath)
			if err := repoScanner.Scan(*noop, projectPath); err != nil {
				Logf(ansi.Red(), "Scanner failed on %s\nError:%v\n", projectPath, err)
				os.Exit(exit.GeneralError)
			}
			Logf(ansi.Green(), "Scan ok: %s", projectPath)
		}
		Log(ansi.Green(), "Scanning successful")
		return
	}
	/*
	 * test all enabled projects
	 */
	if *test {
		for _, projectPath := range projectList {
			Logf(ansi.Yellow(), "Testing %s...", projectPath)
			if err := repoTester.Testing(*noop, projectPath); err != nil {
				Logf(ansi.Red(), "Tester failed on %s\nError:%v\n", projectPath, err)
				os.Exit(exit.GeneralError)
			}
			Logf(ansi.Green(), "Test ok: %s", projectPath)
		}
		Log(ansi.Green(), "Testing successful")
		return
	}
	/*
	 * build all enabled projects
	 */
	if *build {
		for _, projectPath := range projectList {
			Logf(ansi.Yellow(), "Build %s...", projectPath)
			if err := repoBuilder.Build(*noop, projectPath); err != nil {
				Logf(ansi.Red(), "Build failed on %s\nError:%v\n", projectPath, err)
				os.Exit(exit.GeneralError)
			}
			Logf(ansi.Green(), "Build ok: %s", projectPath)
		}
		Log(ansi.Green(), "Build successful")
		return
	}
}
