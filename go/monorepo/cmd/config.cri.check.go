package cmd

import (
	"github.com/spf13/cobra"
)

// configCriCheckCmd represents the 'config command
var configCriCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "- Validate a container runtime configuration",
	Long:  `Validate a container runtime configuration`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	ansi.Red().
	//		Println("'config' must be called with a specific object class to be configured.").
	//		Println("Object classes include 'cri, host, hypervisor, language, platform, project.'").
	//		Reset()
	//},
}

func init() {
	configCriCmd.AddCommand(configCriCheckCmd)
}
