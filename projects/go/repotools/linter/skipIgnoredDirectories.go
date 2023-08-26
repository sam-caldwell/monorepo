package repolinter

import (
	"github.com/sam-caldwell/go/v2/projects/go/misc/words"
	"strings"
)

// skipIgnoredDirectories - skip any 'IgnoredDirectories'
func skipIgnoredDirectories(path string) (skip bool) {
	ignoreCriteria := strings.Split(IgnoredDirectories, words.Comma)
	for _, criteria := range ignoreCriteria {
		skip = skip || strings.Contains(path, criteria)
	}
	return skip
}
