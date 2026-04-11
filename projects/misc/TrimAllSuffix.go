package misc

import "strings"

// TrimAllSuffix removes all instances of the given suffix from the input string.
func TrimAllSuffix(input, suffix string) string {
	for strings.HasSuffix(input, suffix) {
		input = input[:len(input)-len(suffix)]
	}
	return input
}
