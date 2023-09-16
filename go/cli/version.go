package cli

/*
 * cli/version.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the reusable persistent --version flag
 * for cobra cli applications.
 */
import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

// Version - create --noop flag for cobra applications
func Version(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP(words.Version,
		words.EmptyString,
		false,
		"Show current version")
}
