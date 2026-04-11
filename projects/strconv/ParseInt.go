package strconv

import "strconv"

// ParseInt - Wrapper for golang strconv.ParseInt because I am tired of refactoring deprecated stuff golang changes
func ParseInt(s string) (int, error) {
	value, err := strconv.ParseInt(s, 10, 32)
	return int(value), err
}
