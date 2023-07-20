package listprojects

/*
 * projects/repotools/listprojects/SortedLexically.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the SortedLexically() function which will
 * return a list of projects, sorted lexically.
 *
 * See README.md
 */

import (
	keyvalue "github.com/sam-caldwell/go/v2/projects/KeyValue"
	"github.com/sam-caldwell/go/v2/projects/KeyValue/pair"
	"github.com/sam-caldwell/go/v2/projects/repotools/filters"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
)

// SortedLexically - return a list of projects, sorted lexically
func SortedLexically(filter filters.Filter) (list pair.OrderedPair, err error) {
	var projects keyvalue.KeyValue

	projects, err = UnsortedProjects(filter)
	if err != nil {
		return list, err
	}

	err = projects.Walk(func(thisProject string, raw interface{}) error {
		list.Add(thisProject, raw.(projectmanifest.Manifest))
		return nil
	})
	if err == nil {
		list.SortByKey()
	}
	return list, err
}
