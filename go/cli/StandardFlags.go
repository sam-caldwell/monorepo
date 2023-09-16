package cli

import "github.com/spf13/cobra"

/*
 * cli/standardFlags.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the standard flags for
 * cli tools in this repo should adopt, including:
 *     --debug
 *     --version
 *     --nocolor
 *     --noop
 */

func StandardFlags(cmd *cobra.Command) {
	Debug(cmd)
	NoColor(cmd)
	Noop(cmd)
	Version(cmd)
}
