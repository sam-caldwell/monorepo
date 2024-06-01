package simpleArgs

import (
	cliArgs "github.com/sam-caldwell/monorepo/go/misc/cli"
	"os"
)

// QuietMode - return true if the quiet flag is used.
//
//	(c) 2023 Sam Caldwell.  MIT License
func QuietMode() bool {
	for _, arg := range os.Args[1:] {
		if arg == cliArgs.Q || arg == cliArgs.Quiet {
			return true
		}
	}
	return false
}
