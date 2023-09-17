package cli

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

/*
 * cli/IsNoop.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the package-global noop variable
 * and the IsNoop() function used to check if the cli
 * is in a no-operation mode.
 */

// IsNoop - Check if --noop was used and return boolean state
func IsNoop(cmd *cobra.Command) bool {
	return cmd.PersistentFlags().Lookup(words.Noop).Value.String() == words.True
}
