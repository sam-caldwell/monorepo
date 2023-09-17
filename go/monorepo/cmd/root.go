package cmd

/*
 * monorepo/cmd/root.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This is the root of the monorepo command line interface.
 */

import (
	"github.com/sam-caldwell/monorepo/go/cli"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "monorepo",
	Short: "Tooling to manage the monorepo",
	Long: `This command allows simple management of the monorepo from a single point of interaction.

Use this command to create projects, configure the monorepo itself or to build projects, run tests, run
linters or perform security scans, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.EvaluateStandardFlags(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(exit.GeneralError)
	}
}

func init() {
	//Reusable, standardized flags should make app development pretty easy.
	cli.CreateFlagConfigFile(rootCmd, defaultConfigFile())
	cli.CreateStandardFlags(rootCmd)
}
