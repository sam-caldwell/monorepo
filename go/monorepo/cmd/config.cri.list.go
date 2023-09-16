package cmd

import (
	"github.com/sam-caldwell/monorepo/go/cli"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	monorepoCri "github.com/sam-caldwell/monorepo/go/monorepo/cri"
	"github.com/spf13/cobra"
)

// configCriListCmd represents the 'config command
var configCriListCmd = &cobra.Command{
	Use:   "list",
	Short: "- List supported container runtimes",
	Long:  `List supported container runtimes`,
	Run: func(cmd *cobra.Command, args []string) {
		monorepoCri.List(cli.GetFilterList(cmd, words.Opsys))
	},
}

func init() {
	configCriCmd.AddCommand(configCriListCmd)

	cli.FilterByList(
		configCriListCmd,
		words.Opsys,
		cli.GetCurrentOpsys(configCriListCmd),
		"filter by opsys")
}
