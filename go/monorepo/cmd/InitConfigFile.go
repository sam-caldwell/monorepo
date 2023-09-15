package cmd

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/viper"
)

func InitConfigFile() {
	fmt.Println("InitConfigFile")
	cfgFile := rootCmd.Flags().Lookup("config").Value.String()
	if !file.Exists(cfgFile) {
		ansi.CheckErr(rootCmd, fmt.Errorf("config file not found (%s)", cfgFile))
	}
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		ansi.CheckErr(rootCmd, err)
	}
	if rootCmd.PersistentFlags().Lookup(words.Debug).Value.String() == words.True {
		ansi.Debugf(rootCmd, "Using config file: %s\n", viper.ConfigFileUsed())
	}
}
