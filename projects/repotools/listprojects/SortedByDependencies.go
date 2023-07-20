package listprojects

/*
 * projects/repotools/listprojects/SortedByDependencies.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the SortedByDependencies() function which will
 * provide a filtered list of projects sorted by dependency.
 *
 * See README.md
 */

import (
	keyvalue "github.com/sam-caldwell/go/v2/projects/keyvalue"
	"github.com/sam-caldwell/go/v2/projects/keyvalue/pair"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
)

// SortedByDependencies - List projects, in dependency order.
func SortedByDependencies(filter filters.Filter) (list pair.OrderedPair, err error) {
	var projects keyvalue.KeyValue

	projects, err = UnsortedProjects(filter)
	if err != nil {
		return list, err
	}

	err = projects.Walk(func(thisProject string, raw interface{}) error {
		manifest := raw.(projectmanifest.Manifest)
		err := resolveDependencies(&projects, &list, &thisProject, &manifest)
		return err
	})
	return list, err
}
