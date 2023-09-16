package cmd

/*
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 */

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

// configCriCmd represents the 'config cri' command
var configCriCmd = &cobra.Command{
	Use:   "cri",
	Short: "- Manage the monorepo container runtime integration",
	Long:  `This manages any supported container runtime integrations (e.g. docker, kubernetes)`,
}

func init() {
	configCmd.AddCommand(configCriCmd)
	configCriCmd.Flags().BoolP(
		words.NoColor,
		words.EmptyString,
		false,
		words.ConfigCriVerboseUsage)
}
