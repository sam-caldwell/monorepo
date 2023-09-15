package cli

/*
 * cli/cobraFlags/debug.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the reusable persistent --debug flag
 * for cobra cli applications.
 */
import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

var debugFlag bool

// Debug - create --debug flag for cobra applications
func Debug(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP(words.Debug, words.EmptyString, false, "Print verbose output")
}

// SetDebugIfPresent - Set Global Debug flag if present.
func SetDebugIfPresent(cmd *cobra.Command) bool {
	debugFlag = cmd.PersistentFlags().Lookup(words.Debug).Value.String() == words.True
	return debugFlag
}

func init() {
	debugFlag = false
}
