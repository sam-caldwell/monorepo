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
	"github.com/sam-caldwell/go/v2/projects/keyvalue"
	"github.com/sam-caldwell/go/v2/projects/keyvalue/pair"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
	"path/filepath"
)

// SortedByDependencies - List projects, in dependency order.
func SortedByDependencies(filter filters.Filter) (list pair.OrderedPair, err error) {
	var projects keyvalue.KeyValue

	projects, err = UnsortedProjects(filter)
	if err != nil {
		return list, err
	}

	err = projects.Walk(func(parentKey string, parentValue any) error {
		projectRoot := filepath.Dir(parentKey)
		manifest := parentValue.(projectmanifest.Manifest)
		for _, dependency := range manifest.Dependencies {
			if list.Has(filepath.Join(projectRoot, dependency)) {
				continue
			}
			err = projects.Walk(func(depKey string, depValue any) error {
				depManifest := depValue.(projectmanifest.Manifest)
				if depKey != filepath.Join(projectRoot, dependency) {
					return nil
				}
				list.Add(depKey, depManifest)
				return nil
			})
		}
		if !list.Has(parentKey) {
			list.Add(parentKey, parentValue)
		}
		return nil
	})
	if err != nil {
		return list, err
	}
	return list, err
}
