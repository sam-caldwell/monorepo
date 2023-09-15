package cli

/*
 * cli/cobraFlags/version.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the reusable persistent --version flag
 * for cobra cli applications.
 */
import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/version"
	"github.com/spf13/cobra"
)

// Version - create --noop flag for cobra applications
func Version(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP(words.Version,
		words.EmptyString,
		false,
		"Show current version")
}

// ShowVersionIfPresent - Show version if --version is present and exit (0)
func ShowVersionIfPresent(cmd *cobra.Command) (response bool) {
	if cmd.PersistentFlags().Lookup(words.Version).Value.String() == words.True {
		ansi.Println(version.Version)
		response = true
	}
	return response
}
