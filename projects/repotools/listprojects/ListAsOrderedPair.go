package listprojects

import (
	"github.com/sam-caldwell/go/v2/projects/keyvalue"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
)

func ListAsOrderedPair(filter filters.Filter) ([]keyvalue.OrderedPair, error) {
	raw, err := ListProjects(filter)
	return raw.ToOrderedList(true), err
}
