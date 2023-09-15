package cmd

/*
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 */

import (
	"fmt"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var criCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
	},
}

func init() {
	criCmd.AddCommand(criCheckCmd)
}
