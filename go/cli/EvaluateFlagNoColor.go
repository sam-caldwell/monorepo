package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

/*
 * cli/EvaluateFlagNoColor.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the reusable persistent --debug flag
 * for cobra cli applications.
 */

// EvaluateFlagNoColor - Set Ansi NoColor flag if present
func EvaluateFlagNoColor(cmd *cobra.Command) {
	if cmd.PersistentFlags().Lookup(words.NoColor).Value.String() == words.True {
		ansi.Disable()
	} else {
		ansi.Enable()
	}

}
