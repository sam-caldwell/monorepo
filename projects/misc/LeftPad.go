package misc

import (
	"fmt"
	"strings"
)

// LeftPad - pad the left side of string (s) with char (c) for total output size (sz) chars
func LeftPad(c rune, s string, sz int) string {
	if padSz := sz - len(s); padSz > 0 {
		return fmt.Sprintf("%s%s", strings.Repeat(string(c), padSz), s)
	}
	return s
}
