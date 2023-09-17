package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
)

/*
 * cli/EvaluateFlagNoop.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the reusable persistent --debug flag
 * for cobra cli applications.
 */

// EvaluateFlagNoop - Set Ansi CreateFlagNoop flag if present
func EvaluateFlagNoop(cmd *cobra.Command) {
	titleCase := func(b bool) string {
		c := cases.Title(language.Und)
		return c.String(strconv.FormatBool(b))
	}
	ansi.Debugf(
		"CreateFlagNoop Mode  : %s\n",
		titleCase(cmd.PersistentFlags().Lookup(words.Noop).Value.String() == words.True))
}
