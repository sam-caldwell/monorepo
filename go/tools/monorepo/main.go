package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	monorepo "github.com/sam-caldwell/monorepo/go/tools/monorepo/lib"
	"path/filepath"
	"strings"
)

func main() {

	var err error

	rootPath := filepath.Dir(directory.GetCurrent())
	if !directory.Exists(filepath.Join(rootPath, "monorepo")) {
		ansi.Red().
			Printf("This must be run from the root of the monorepo").
			LF().Fatal(exit.GeneralError)
	}

	commands, debug := monorepo.GetArgs()

	Monorepo := monorepo.Monorepo{
		Debug: *debug,
		Root:  rootPath,
	}
	monorepo.PrintHeader("Monorepo command")
	ansi.Cyan().Println("Discovering project manifests")
	if err := Monorepo.GetManifestList(); err != nil {
		ansi.Red().
			Printf("Error discovering manifests\n%s", err).
			Reset().
			Fatal(exit.GeneralError)
	}
	ansi.Green().Printf("Discovered %d manifests\n", Monorepo.ManifestCount()).LF().Reset()

	for _, command := range commands {
		monorepo.PrintHeader(convert.Capitalize(command))
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
		monorepo.PrintFooter(&command, err)
	}
}
