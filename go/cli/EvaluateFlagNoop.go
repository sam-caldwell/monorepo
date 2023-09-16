package cli

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

/*
 * cli/EvaluateFlagNoop.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the reusable persistent --debug flag
 * for cobra cli applications.
 */

// EvaluateFlagNoop - Set Ansi Noop flag if present
func EvaluateFlagNoop(cmd *cobra.Command) {
	if cmd.PersistentFlags().Lookup(words.Noop).Value.String() == words.True {
		noop = true
	} else {
		noop = false
	}
}
