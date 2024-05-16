package convert

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

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
