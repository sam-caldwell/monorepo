package misc

import "strings"

func TrimSuffixCount(input, suffix string, count int) string {
	for i := 0; i < count; i++ {
		input = strings.TrimSuffix(input, suffix)
	}
	return input
}
