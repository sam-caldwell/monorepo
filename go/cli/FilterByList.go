package cli

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
	"strings"
)

// FilterByList - Filter by comma-delimited list against a given configPath
func FilterByList(cmd *cobra.Command, configPath string, filterDefault []string, helpString string) {
	defaultValue := strings.Join(filterDefault, words.Comma)
	cmd.Flags().String(configPath, defaultValue, helpString)
}
