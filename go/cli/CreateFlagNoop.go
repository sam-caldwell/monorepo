package cli

/*
 * cli/CreateFlagNoop.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the reusable persistent --noop flag
 * for cobra cli applications.
 */
import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

// CreateFlagNoop - create --noop flag for cobra applications
func CreateFlagNoop(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP(words.Noop,
		words.EmptyString,
		false,
		"execute command without performing any persistent changes.")
}
