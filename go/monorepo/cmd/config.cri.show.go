package cmd

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	monorepoCri "github.com/sam-caldwell/monorepo/go/monorepo/cri"
	"github.com/spf13/cobra"
)

// configCriShowCmd represents the 'config command
var configCriShowCmd = &cobra.Command{
	Use:   "show",
	Short: "- Show a specified container runtime configuration",
	Long:  `Show a specified container runtime configuration`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ansi.Blue().Println("config cri show <name>")
		name := args[0]
		ansi.Blue().Printf("config cri show %s", name).LF().Reset()
		monorepoCri.Show(name)
	},
}

func init() {
	configCriCmd.AddCommand(configCriShowCmd)
}
