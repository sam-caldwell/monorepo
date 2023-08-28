package monorepo

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/sets/simple"
	"strings"
)

func GetSupportedArch() []string {
	const ArchColumn = 1
	var set simple.Set
	for _, record := range strings.Split(SupportedPlatforms, ",") {
		parts := strings.Split(record, ":")
		if err := set.Add(parts[ArchColumn]); err != nil {
			panic(err)
		}
	}
	return set.ListString(true)
}
