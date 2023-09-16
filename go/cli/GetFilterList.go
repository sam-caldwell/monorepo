package cli

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
	"strings"
)

// GetFilterList - Get the filter list (--<option>) string and return the parsed comma-delimited list.
func GetFilterList(cmd *cobra.Command, flagName string) (filterSet []string) {
	//Use a map to create a unique opsys set.
	uniqueSet := Set{}
	for _, f := range strings.Split(cmd.Flags().Lookup(flagName).Value.String(), words.Comma) {
		uniqueSet.Add(f)
		switch f {
		case words.Darwin:
			uniqueSet.Add(words.Macos)
		case words.Macos:
			uniqueSet.Add(words.Darwin)
		case words.Ubuntu, words.Debian, words.Redhat:
			uniqueSet.Add(words.Linux)
		case words.Windows2K, words.WindowsXp, words.WindowsVista, words.WindowsServer2K3, words.Windows7,
			words.Windows8, words.Windows10, words.Windows:
			uniqueSet.Add(words.Windows)
		}
	}
	return uniqueSet.ToArray()
}
