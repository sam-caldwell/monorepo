package misc

import (
	"fmt"
	"strings"
)

// FlattenStringMap - Flatten the map to a string
func FlattenStringMap(data map[string]string) string {
	var lines []string

	keyWidth := MaxKeyWidth(data)

	// Format each key-value pair with uniform columns
	for key, value := range data {
		line := fmt.Sprintf("%-*s : %s", keyWidth, key, value)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
