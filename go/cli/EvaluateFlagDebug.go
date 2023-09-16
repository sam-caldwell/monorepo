package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

/*
 * cli/cobraFlags/debug.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the reusable persistent --debug flag
 * for cobra cli applications.
 */

// EvaluateFlagDebug - Set Ansi Debug flag if present
func EvaluateFlagDebug(cmd *cobra.Command) {
	if cmd.PersistentFlags().Lookup(words.Debug).Value.String() == words.True {
		ansi.EnableDebug()
	} else {
		ansi.DisableDebug()
	}
}
