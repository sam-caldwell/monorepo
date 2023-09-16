package cli

import (
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/viper"
)

// GetCriList - Compare the filter list (from cli) to the cri list from the config file and return the
// names of container runtimes which match the filter.
func GetCriList(filterSet []string) []string {
	platformNames := Set{}
	containerRuntimes := viper.Get(words.ContainerRuntimes).([]interface{})
	for _, r := range containerRuntimes {
		row := r.(map[string]interface{})
		platformList := convert.InterfacesToStrings(row["platforms"].([]any))
		if e := row["enabled"].(bool); e {
			for _, platform := range platformList {
				for _, filter := range filterSet {
					if filter == platform {
						platformNames.Add(row["name"].(string))
					}
				}
			}
		}
	}
	return platformNames.ToArray()
}
