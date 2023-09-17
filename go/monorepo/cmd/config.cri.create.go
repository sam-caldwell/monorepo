package cmd

import (
	"github.com/spf13/cobra"
)

// configCriCreateCmd represents the 'config command
var configCriCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "- Add a new supported container runtime configuration",
	Long:  `Add a new supported container runtime configuration`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	ansi.Red().
	//		Println("'config' must be called with a specific object class to be configured.").
	//		Println("Object classes include 'cri, host, hypervisor, language, platform, project.'").
	//		Reset()
	//},
}

func init() {
	configCriCmd.AddCommand(configCriCreateCmd)
}
