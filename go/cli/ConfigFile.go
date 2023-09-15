package cli

/*
 * cli/cobraFlags/ConfigFile.go
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

// ConfigFile - create --config persistent flag (cobra) and load any config file
func ConfigFile(cmd *cobra.Command, defaultConfigFile string) {
	cobra.OnInitialize(initConfigFile(cmd))
	cmd.PersistentFlags().String(words.Config, defaultConfigFile, "config file")
}
