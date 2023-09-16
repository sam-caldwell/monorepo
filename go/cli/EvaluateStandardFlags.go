package cli

/*
 * cli/EvaluateStandardFlags.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file evaluates the standard flags for
 * cli tools in this repo should adopt, including:
 *     --debug
 *     --version
 *     --nocolor
 *     --noop
 */

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strconv"
)

// EvaluateStandardFlags - Evaluate the standard tags (if present) and display any active debug, noop modes.
func EvaluateStandardFlags(cmd *cobra.Command) {
	titleCase := func(b bool) string {
		c := cases.Title(language.Und)
		return c.String(strconv.FormatBool(b))
	}
	EvaluateFlagNoColor(cmd)
	EvaluateFlagDebug(cmd)
	EvaluateFlagNoop(cmd)
	ansi.Debugf("Debug Mode : Enabled\n")
	ansi.Debugf("Noop Mode  : %s\n", titleCase(noop))
	if !ShowVersionIfPresent(cmd) { // show version and terminate.
		ansi.Println(cmd.UsageString()).Reset()
		os.Exit(exit.GeneralError)
	}
}
