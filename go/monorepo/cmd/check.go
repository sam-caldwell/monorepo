package cmd

/*
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 */

import (
	"github.com/sam-caldwell/monorepo/go/ansi"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Perform config validation",
	Long:  `This command is used to validate a specific part of the monorepo configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		ansi.Red().
			Println("check called without context").
			LF().
			Println("ToDo: Call all checks when no context is provided.").
			Reset()

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
