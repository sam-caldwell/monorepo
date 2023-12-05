package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	monorepo "github.com/sam-caldwell/monorepo/go/tools/monorepo/lib"
	"github.com/sam-caldwell/monorepo/go/version"
	"path/filepath"
	"strings"
)

func main() {

	command := monorepo.GetCommand()

	cmdVersion := flag.Bool("version", false, "show version")

	class := flag.String("class", "all", "specify a project class (e.g. go,containers, js, cpp")

	project := flag.String("project", "all", "specify a group of projects to build")

	debug := flag.Bool("debug", false, "print debug messages")

	flag.Parse()
	if *cmdVersion {
		version.Show()
	}

	if *debug {
		ansi.Blue().
			Println("Debug:").
			Printf("\tcommand: %s\n", *command).
			Printf("\tclass:   %s\n", *class).
			Printf("\tproject: %s\n", *project).
			LF().
			Reset()
	}

	switch *command {
	case "build":
		if err := monorepo.Build(class, project); err != nil {
			ansi.Red().Printf("error on build: %v", err).LF().Fatal(exit.GeneralError)
		}
		ansi.Green().Println("Build Successful").Reset()

	case "clean":
		if err := monorepo.Clean(class, project); err != nil {
			ansi.Red().Printf("error on build: %v", err).LF().Fatal(exit.GeneralError)
		}
		ansi.Green().Println("Clean Successful").Reset()

	case "test":
		if err := monorepo.Test(class, project); err != nil {
			ansi.Red().Printf("error on build: %v", err).LF().Fatal(exit.GeneralError)
		}
		ansi.Green().Println("Test Successful").Reset()

	case "list":
		ansi.Green().
			Println("List of projects").
			Println(strings.Repeat("=", 120)).
			LF().Reset()

		list, err := monorepo.GetProjectList(class, project)
		if err != nil {
			ansi.Red().Printf("Error:%v", err).LF().Reset().Fatal(exit.InvalidInput)
		}
		for class, projectList := range list {
			ansi.Green().Printf(" class: %s\n", filepath.Base(class))
			for _, project := range projectList {
				projectName := strings.TrimPrefix(project, directory.GetCurrent()+"/"+filepath.Base(class)+"/")
				ansi.Blue().Printf("\t%-42s\n", projectName).Reset()
			}
		}
		ansi.Green().Println("Ok").Reset()

	case "help":
		ansi.Blue().LF().Println("Usage:").LF().
			Println("Command (build, clean,list,test)").LF().
			Println("Options:").LF()
		flag.PrintDefaults()
		ansi.LF().Reset()

	default:

		ansi.Red().
			Println("Missing or unsupported command").
			LF().Reset().Fatal(exit.InvalidInput)

	}

}
