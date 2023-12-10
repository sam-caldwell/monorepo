package main

/*
 * monorepo/sqlTrackerDb.go
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This is the main program for the monorepo command
 * This is the start of all the automation goodness.
 */

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	monorepo "github.com/sam-caldwell/monorepo/go/tools/monorepo/lib"
	"strings"
	"time"
)

func main() {

	var err error

	rootPath := directory.GetCurrent()
	if !directory.Exists(rootPath) {
		ansi.Red().
			Printf("This must be run from the root of the monorepo").
			LF().Fatal(exit.GeneralError)
	}

	commands, debug := monorepo.GetArgs()

	Monorepo := monorepo.Monorepo{
		StartTime: time.Now(),
		Debug:     *debug,
		Root:      rootPath,
	}

	Monorepo.PrintHeader("Monorepo command")
	Monorepo.LoadManifests()
	ansi.Cyan().Printf("%d records loaded", Monorepo.ManifestCount()).LF().Reset()
	Monorepo.SortByClassAndProject()

	for _, command := range commands {
		Monorepo.PrintHeader(convert.Capitalize(command))
		switch strings.ToLower(strings.TrimSpace(command)) {
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
		Monorepo.PrintFooter(&command, err)
	}
}
