package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/exit"
	monorepo "github.com/sam-caldwell/monorepo/go/tools/monorepo/lib"
)

func main() {

	var err error

	command, class, project, debug := monorepo.GetArgs()

	Monorepo := monorepo.Monorepo{
		Debug: *debug,
	}

	switch *command {

	case "build":
		err = Monorepo.Build(class, project)

	case "clean":
		err = Monorepo.Clean(class, project)

	case "help":
		ansi.Blue().LF().Println("Usage:").LF().
			Println("Command (build, clean,list,test)").LF().
			Println("Options:").LF()
		flag.PrintDefaults()
		ansi.LF().Reset()

	case "list":
		err = Monorepo.List(class, project)

	case "test":
		err = Monorepo.Test(class, project)

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
		Printf("%s Successful", convert.Capitalizep(command)).
		LF().Reset()

}
