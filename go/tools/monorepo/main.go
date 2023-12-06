package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	monorepo "github.com/sam-caldwell/monorepo/go/tools/monorepo/lib"
	"path/filepath"
)

func main() {

	var err error

	command, debug := monorepo.GetArgs()

	rootPath := filepath.Dir(directory.GetCurrent())
	if !directory.Exists(filepath.Join(rootPath, "monorepo")) {
		ansi.Red().
			Printf("This must be run from the root of the monorepo").
			LF().Fatal(exit.GeneralError)
	}

	Monorepo := monorepo.Monorepo{
		Debug: *debug,
		Root:  rootPath,
	}

	switch *command {

	case "build":
		err = Monorepo.Build()

	case "clean":
		err = Monorepo.Clean()

	case "help":
		ansi.Blue().LF().Println("Usage:").LF().
			Println("Command (build, clean,list,test)").LF().
			Println("Options:").LF()
		flag.PrintDefaults()
		ansi.LF().Reset()

	case "list":
		err = Monorepo.List()

	case "test":
		err = Monorepo.Test()

	default:
		ansi.Red().
			Println("Missing or unsupported command").
			LF().Reset().Fatal(exit.InvalidInput)
	}

	if err != nil {
		ansi.Red().
			Printf("error on %s: %v", convert.Capitalizep(command), err).
			LF().Fatal(exit.GeneralError)
	}

	ansi.Green().LF().
		Printf("%s Successful\n", convert.Capitalizep(command)).
		LF().Reset()

}
