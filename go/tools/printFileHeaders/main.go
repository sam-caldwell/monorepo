package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/binaryAnalysis"
	executables "github.com/sam-caldwell/monorepo/go/binaryAnalysis/Headers"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
)

func main() {
	fileName := flag.String("file", words.EmptyString, "the file name to analyze")
	debug := flag.Bool(words.Debug, false, "print debug messages")
	flag.Parse()

	if !file.Exists(*fileName) {
		ansi.Red().Println("file not found").Reset().Fatal(exit.NotFound)
	}
	fh, err := os.Open(*fileName)
	if err != nil {
		ansi.Red().Printf("cannot open file: %s", err).LF().Reset().Fatal(exit.FailedToOpenFile)
	}

	format, err := binaryAnalysis.GetFileType(fh, *debug)
	if err != nil {
		ansi.Red().Printf("Error detecting format: %s", err).Reset().Fatal(exit.GeneralError)
	}

	switch format {
	case binaryAnalysis.Elf32:
		var header executables.Elf32
		if err := header.Load(fh); err != nil {
			ansi.Red().Printf("Error: %s", err).Reset().Fatal(exit.GeneralError)
		}

	case binaryAnalysis.Elf64:
		var header executables.Elf64
		if err := header.Load(fh); err != nil {
			ansi.Red().Printf("Error: %s", err).Reset().Fatal(exit.GeneralError)
		}
	default:
		ansi.Red().Println("unsupported file format").Reset().Fatal(exit.GeneralError)
	}
}
