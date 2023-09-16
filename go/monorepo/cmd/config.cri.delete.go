package cmd

/*
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 */

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/spf13/cobra"
)

// configCriDeleteCmd represents the delete command
var configCriDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "- Delete a container runtime configuration",
	Long:  `Delete a container runtime configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		ansi.Blue().Println("delete config <name>").Reset()
	},
}

func init() {
	configCriCmd.AddCommand(configCriDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	configCriDeleteCmd.Flags().String("name", "", "Specify the project name to delete (required)")
}
