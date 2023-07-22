package main

import "github.com/sam-caldwell/go/v2/projects/exit"

const (
	commandUsage = `
`
)

func main() {

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

}
