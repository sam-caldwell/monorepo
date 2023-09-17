package cmd

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/cli"
	monorepoCri "github.com/sam-caldwell/monorepo/go/monorepo/cri"
	"github.com/spf13/cobra"
)

// configCriEnableCmd represents the 'config command
var configCriEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "- Enable a specified container runtime configuration",
	Long: `
Enable a specified container runtime configuration
`,
	Args: cobra.ExactArgs(1), //name
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if cli.IsNoop(cmd) {
			ansi.Yellow().Printf("CRI config (%s) enable (noop)\n", name).Reset()
			return
		}
		if err := monorepoCri.Enable(name); err != nil {
			ansi.Red().Printf("CRI config (%s) enable failed\n", name).Reset()
		}
		ansi.Green().Printf("CRI config (%s) enabled\n", name).Reset()
		cli.WriteConfig(cmd)
	},
}

func init() {
	configCriCmd.AddCommand(configCriEnableCmd)
	cli.CreateStandardFlags(configCriEnableCmd)
}
