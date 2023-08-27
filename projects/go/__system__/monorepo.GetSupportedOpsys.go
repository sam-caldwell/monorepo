package monorepo

import (
	"github.com/sam-caldwell/go/v2/projects/go/sets/simple"
	"strings"
)

func GetSupportedOpsys() (out []string) {
	const OpsysColumn = 0
	var set simple.Set
	for _, record := range strings.Split(SupportedPlatforms, ",") {
		parts := strings.Split(record, ":")
		if err := set.Add(parts[OpsysColumn]); err != nil {
			panic(err)
		}
	}
	return set.ListString(true)
}
