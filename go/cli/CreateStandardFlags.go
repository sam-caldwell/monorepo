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

// CreateStandardFlags - Create the standard flags (--debug, --nocolor, --noop and --version)
func CreateStandardFlags(cmd *cobra.Command) {
	CreateFlagDebug(cmd)
	CreateFlagNoColor(cmd)
	CreateFlagNoop(cmd)
	Version(cmd)
}
