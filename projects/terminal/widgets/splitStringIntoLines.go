package terminal

import "strings"

// SplitStringIntoLines -Given a string, split the line into a list of lines no longer than n-characters.
func SplitStringIntoLines(s string, n int) (results []string) {
	// Check if n is valid
	if n <= 0 {
		panic("n must be greater than 0")
	}
	s = strings.TrimSpace(s)
	for i := 0; i < len(s); i += n {
		var subStr string
		if (i + n) > len(s)-1 {
			subStr = s[i:]
		} else {
			subStr = s[i : i+n]
			if len(subStr) < n {
				subStr = s[i:]
			}
		}
		results = append(results, subStr)
	}

	return results
}
