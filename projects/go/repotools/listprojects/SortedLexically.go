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
	"github.com/sam-caldwell/go/v2/projects/go/keyvalue"
	"github.com/sam-caldwell/go/v2/projects/go/keyvalue/pair"
	"github.com/sam-caldwell/go/v2/projects/go/repotools/filters"
)

// SortedLexically - return a list of projects, sorted lexically
func SortedLexically(filter filters.Filter) (list pair.OrderedPair, err error) {
	var projects keyvalue.KeyValue

	projects, err = UnsortedProjects(filter)
	if err != nil {
		return list, err
	}
	list = projects.ToOrderedPairList(true)
	return list, err
}
