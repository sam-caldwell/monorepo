package cli

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func WriteConfig(cmd *cobra.Command) {
	if err := viper.WriteConfig(); err != nil {
		ansi.CheckErr(cmd, fmt.Sprintf("cli.WriteConfig() failed: %v", err))
	}
}
