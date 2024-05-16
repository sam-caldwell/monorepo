package tags

import (
	"fmt"
	"strings"
)

// String - return a delimited list of key:value pairs
func (tags *Tags) String(kvDelimiter, termDelimiter string) (output string) {
	var result []string
	for key, value := range tags.data {
		pair := fmt.Sprintf("%s%s%s", key, kvDelimiter, value)
		result = append(result, pair)
	}
	return strings.Join(result, termDelimiter)
}
