package convert

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

// MapToString - Given a map of generic types as a string.
//
//	(c) 2023 Sam Caldwell.  MIT License
func MapToString[K string, V any](tags map[K]V) string {
	var tagPairs []string

	if tags == nil {
		return words.EmptyString
	}

	for k, v := range tags {
		tagPairs = append(tagPairs, fmt.Sprintf("%v=%v", k, v))
	}

	return strings.Join(tagPairs, "&")

}
