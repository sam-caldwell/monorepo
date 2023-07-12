package monorepo

import (
	"github.com/sam-caldwell/go/v2/projects/sets/simple"
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
