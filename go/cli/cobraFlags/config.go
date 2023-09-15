package cobraFlags

/*
 * cli/cobraFlags/config.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides
 */
import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ConfigFlag(cmd *cobra.Command, defaultConfigFile string) {
	cobra.OnInitialize(InitConfigFile(cmd))
	cmd.PersistentFlags().String(words.Config, defaultConfigFile, "config file")
}

func InitConfigFile(cmd *cobra.Command) func() {
	return func() {
		fmt.Println("InitConfigFile")
		cfgFile := cmd.Flags().Lookup("config").Value.String()
		if !file.Exists(cfgFile) {
			ansi.CheckErr(cmd, fmt.Errorf("config file not found (%s)", cfgFile))
		}
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			ansi.CheckErr(cmd, err)
		}
		if cmd.PersistentFlags().Lookup(words.Debug).Value.String() == words.True {
			ansi.Debugf(cmd, "Using config file: %s\n", viper.ConfigFileUsed())
		}
	}
}
