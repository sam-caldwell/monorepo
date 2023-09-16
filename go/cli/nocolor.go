package cli

/*
 * cli/cobraFlags/nocolor.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the --nocolor persistent flag
 * for cobra cli applications.
 */
import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

// NoColor - turn off ANSI color output for cobra applications
func NoColor(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP(words.NoColor, words.EmptyString, false, "Print verbose output")
}
