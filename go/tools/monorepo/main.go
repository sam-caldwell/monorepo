package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	monorepo "github.com/sam-caldwell/monorepo/go/tools/monorepo/lib"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
)

func main() {
	var command string
	if len(os.Args) > 2 {
		command = os.Args[1] //build,clean,test
		if len(os.Args) > 2 {
			os.Args = os.Args[1:]
		}
	}
	cmdVersion := flag.Bool("version", false, "show version")
	project := flag.String("project", "all", "specify a single project to manipulate.")
	flag.Parse()
	if *cmdVersion {
		version.Show()
	}
	ansi.Blue().Printf("command: %s", command).LF().Reset()
	ansi.Blue().Printf("project: %s", *project).LF().Reset()
	switch command {
	case "build":
		if err := monorepo.Build(project); err != nil {
			ansi.Red().Printf("error on build: %v", err).LF().Fatal(exit.GeneralError)
		}
		ansi.Green().Println("Build Successful").Reset()
	case "clean":
		if err := monorepo.Clean(project); err != nil {
			ansi.Red().Printf("error on build: %v", err).LF().Fatal(exit.GeneralError)
		}
		ansi.Green().Println("Clean Successful").Reset()
	case "test":
		if err := monorepo.Test(project); err != nil {
			ansi.Red().Printf("error on build: %v", err).LF().Fatal(exit.GeneralError)
		}
		ansi.Green().Println("Test Successful").Reset()
	default:
		ansi.Red().Println("Missing or unsupported command").LF().Fatal(exit.InvalidInput)
	}
}
