package misc

import (
	"fmt"
	"strings"
)

// FlattenStringMap - Flatten the map to a string
func FlattenStringMap(data map[string]string) (output string) {
	var lines []string
	for key, value := range data {
		lines = append(lines, fmt.Sprintf("%v:%v", key, value))
	}
	return strings.Join(lines, "\n")
}
