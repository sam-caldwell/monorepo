package simpleArgs

import (
	cliArgs "github.com/sam-caldwell/monorepo/go/misc/cli"
	"os"
)

// QuietMode - return true if the quiet flag is used.
func QuietMode() bool {
	for _, arg := range os.Args[1:] {
		if arg == cliArgs.Q || arg == cliArgs.Quiet {
			return true
		}
	}
	return false
}
