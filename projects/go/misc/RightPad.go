package misc

// RightPad - pad the right side of string (s) with char (c) for total output size (sz) chars
func RightPad(c rune, s string, sz int) string {
	return s + RepeatChars(string(c), sz-len(s))
}
