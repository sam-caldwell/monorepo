package cmd

/*
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/cli/cobraFlags"
	"github.com/sam-caldwell/monorepo/go/exit"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "monorepo",
	Short: "Tooling to manage the monorepo",
	Long: `This command allows simple management of the monorepo from a single point of interaction.

Use this command to create projects, configure the monorepo itself or to build projects, run tests, run
linters or perform security scans, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
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
	defaultConfigFile := func() string {
		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			cobra.CheckErr(err)
		}
		return filepath.Join(homeDirectory, "git", "monorepo", "monorepo.yaml")
	}
	rootCmd.PersistentFlags().BoolP("nocolor", "c", false, "Print output without color")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Print verbose output")
	rootCmd.PersistentFlags().BoolP("noop", "s", false, "Execute without persistent change")

	cobraFlags.ConfigFlag(rootCmd, defaultConfigFile())
	rootCmd.Flags().BoolP("version", "v", false, "Show current version")
}
