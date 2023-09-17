package cli

/*
 * cli/ShowVersionIfPresent.go
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

// ShowVersionIfPresent - Show version if --version is present and exit (0)
func ShowVersionIfPresent(cmd *cobra.Command) (response bool) {
	if cmd.Flags().Lookup(words.Version).Value.String() == words.True {
		ansi.Println(version.Version)
		response = true
	}
	return response
}