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
	"os"
)

// EvaluateStandardFlags - Evaluate the standard tags (if present) and display any active debug, noop modes.
func EvaluateStandardFlags(cmd *cobra.Command) {
	EvaluateFlagNoColor(cmd)
	EvaluateFlagDebug(cmd)
	EvaluateFlagNoop(cmd)
	if !ShowVersionIfPresent(cmd) { // show version and terminate.
		ansi.Println(cmd.UsageString()).Reset()
		os.Exit(exit.GeneralError)
	}
}
