package main

import (
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/monorepo/args"
)

func main() {
	exit.IfHelpRequested(args.CommandUsage)
	exit.IfVersionRequested()
	var arg args.Arguments
}
