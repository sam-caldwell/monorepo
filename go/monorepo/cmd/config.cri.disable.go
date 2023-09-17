package cmd

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/cli"
	monorepoCri "github.com/sam-caldwell/monorepo/go/monorepo/cri"
	"github.com/spf13/cobra"
)

// configCriDisableCmd represents the 'config command
var configCriDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "- Disable a specified container runtime configuration",
	Long: `
Disable a specified container runtime configuration
`,
	Args: cobra.ExactArgs(1), //name
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if cli.IsNoop() {
			ansi.Yellow().Printf("CRI config (%s) disable (noop)\n", name).Reset()
			return
		}
		ansi.Debugf("cri name: %s", name)
		if err := monorepoCri.Disable(name); err != nil {
			ansi.Red().Printf("CRI config (%s) disable failed\n", name).Reset()
		}
		ansi.Green().Printf("CRI config (%s) disabled\n", name).Reset()

		cli.WriteConfig(cmd)
	},
}

func init() {
	configCriCmd.AddCommand(configCriDisableCmd)
}
