package misc

import "fmt"

// LeftPad - pad the left side of string (s) with char (c) for total output size (sz) chars
func LeftPad(c rune, s string, sz int) string {
	return fmt.Sprintf("%*s", sz-len(s), string(c))
}
