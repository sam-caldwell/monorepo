package cli

/*
 * cli/CreateFlagDebug.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the reusable persistent --debug flag
 * for cobra cli applications.
 */
import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

// CreateFlagDebug - create --debug flag for cobra applications
func CreateFlagDebug(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP(words.Debug, words.EmptyString, false, "Print verbose output")
}
