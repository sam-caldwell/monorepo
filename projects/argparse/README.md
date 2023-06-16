argparse
===========

## Description

A python-approximate golang commandline argument parser. The goal is to create a Command-line argument
parser similar to python argparse. Obviously, golang will have some differences.

## Usage

See [example](example/main.go).
```golang
package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse"
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/types"
	"log"
)

// main - Demonstration program

func main() {
	log.Println("starting example...")
	namespace := func() []string {
		// We only need argparse to parse then reduce the arguments...
		var args argparse.Arguments
		return args.
			ProgramName().
			Copyright(2023, "Sam Caldwell", "mail@samcaldwell.net>").
			Preamble("This is a description of our Arguments Object.").
			Postscript("This is the postfix string after help is dumped.").
			Add("all", "-a", "--all", types.Flag, true, false, "All flag").
			Add("number", "-n", "--number", types.Integer, true, 1337, "set number").
			Add("bool", "-b", "--bool", types.Boolean, true, false, "set boolean").
			ExitOnError().
			Parse().
			Reduce()
	}()
	// Afterward, we can deal with just the parsed information...
	fmt.Println(namespace.GetValue("all"))
}

```

