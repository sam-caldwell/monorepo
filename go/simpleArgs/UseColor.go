package simpleArgs

import (
	cliArgs "github.com/sam-caldwell/monorepo/go/misc/cli"
	"os"
)

// UseColor - Return true if the color flag is used
//
//	(c) 2023 Sam Caldwell.  MIT License
func UseColor() bool {
	for _, arg := range os.Args[1:] {
		if arg == cliArgs.Color {
			return true
		}
	}
	return false
}
