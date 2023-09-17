package cmd

/*
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 */

import (
	"github.com/sam-caldwell/monorepo/go/cli"
	"github.com/spf13/cobra"
)

// configCmd represents the 'config' command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Edit the monorepo configuration",
	Long: `This command creates/edits/deletes, lists and shows configuration objects for the monorepo.

For example, this command could be used to--
 - create a project or a build host;
 - add a new language to the monorepo supported languages;
 - configure a supported platform (opsys, CPU, Container Runtime, Hypervisor, etc); or
 - list objects configured for a specific object class (project, host, cri, hypervisor, language).
 - print the current configuration fo a specific object.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.EvaluateStandardFlags(cmd)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	cli.CreateStandardFlags(configCmd)
}
