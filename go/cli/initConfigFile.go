package cli

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*
 * cli/cobraFlags/initConfigFile.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides a function to initialize the config file
 * for a cobra command-line program.
 */

// initConfigFile - load the config file
func initConfigFile(cmd *cobra.Command) func() {
	return func() {
		cfgFile := cmd.Flags().Lookup(words.Config).Value.String()
		if !file.Exists(cfgFile) {
			ansi.CheckErr(cmd, fmt.Errorf("file not found (%s)", cfgFile))
		}
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType(words.Yaml)
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			ansi.CheckErr(cmd, err)
		}
		ansi.Debugf(cmd, "Using config file: %s\n", viper.ConfigFileUsed())
	}
}
