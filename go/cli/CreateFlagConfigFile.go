package cli

/*
 * cli/CreateFlagConfigFile.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file allows cobra command-line applications
 * to quickly define --config flags for the
 * root command without reinventing wheels each time.
 */
import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

// CreateFlagConfigFile - create --config persistent flag (cobra) and load any config file
func CreateFlagConfigFile(cmd *cobra.Command, defaultConfigFile string) {
	cobra.OnInitialize(initConfigFile(cmd))
	cmd.PersistentFlags().String(words.Config, defaultConfigFile, "config file")
}
