package main

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/binaryAnalysis/Disassember"
	"github.com/sam-caldwell/monorepo/go/exit"
	argument "github.com/sam-caldwell/monorepo/go/tools/disassembler/args"
)

const usage = `
    disassembler -h | --help

    disassembler --in <file> --out <file> --method <linear|recursive> --arch <amd64,arm64>
`

func main() {
	ansi.Blue().Println("Starting").Reset()
	var args argument.Data
	//var metadata executables.Headers
	args.Parse(usage)

	err := Disassember.Disassemble(args.Debug(), args.SourceFile(), args.OutFile(), args.Method(), args.Arch(), args.Mode())
	if err != nil {
		ansi.Red().Println(err.Error()).Reset().Fatal(exit.GeneralError)
	}

	ansi.Green().Println("Disassembly Complete").Reset().Fatal(exit.Success)

}
