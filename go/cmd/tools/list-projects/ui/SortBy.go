package ui

import (
	"os"
	"strings"
)

type sortOrder byte

const (
	SortLexically = iota
	SortByDependency
)

func SortBy() sortOrder {
	for _, arg := range os.Args {
		if strings.TrimSpace(strings.TrimPrefix(arg, "-")) == "byDependency" {
			return SortByDependency
		}
	}
	return SortLexically
}
