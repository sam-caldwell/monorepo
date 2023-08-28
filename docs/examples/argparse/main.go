package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse"
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"log"
)

// main - Demonstration program
func main() {
	log.Println("starting example...")
	var args argparse.Arguments
	args.
		ProgramName().
		Copyright(2023, "Sam Caldwell", "mail@samcaldwell.net>").
		Preamble("This is a description of our Arguments Object.").
		Postscript("This is the postfix string after help is dumped.").
		Add("all", "-a", "--all", types.Flag, true, false, "All flag").
		Add("number", "-n", "--number", types.Integer, true, 1337, "set number").
		Add("bool", "-b", "--bool", types.Boolean, true, false, "set boolean").
		ExitOnError().
		Parse()

	fmt.Println(args.Help())

}
