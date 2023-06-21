package misc

import "unicode"

// IsStringCapitalized - return bool if first character of string is capitalized
func IsStringCapitalized(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}
