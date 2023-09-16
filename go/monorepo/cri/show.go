package monorepoCri

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/cli"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/viper"
)

func Show(name string) {
	containerRuntimes := viper.Get(words.ContainerRuntimes).([]interface{})
	criSet := cli.Set{}
	for _, r := range containerRuntimes {
		row := r.(map[string]interface{})
		if row[words.Name].(string) == name {
			platformList := convert.InterfacesToStrings(row["platforms"].([]any))
			criSet.Add(
				fmt.Sprintf("- name: %s\n  enabled: %t\n  platforms: %v\n",
					name, row[words.Enabled].(bool), platformList))
		}
	}
	ansi.Blue().
		Line(words.Hyphen, 40).
		Printf("Container Runtime Configuration (count: %d)\n", len(criSet)).
		Line(words.Hyphen, 40)
	for k, _ := range criSet {
		ansi.Printf("%s", k).LF()
	}
	ansi.Line(words.EqualSign, 40).Reset()
}
