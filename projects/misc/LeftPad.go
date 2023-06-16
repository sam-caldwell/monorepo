package misc

// LeftPad - pad the left side of string (s) with char (c) for total output size (sz) chars
func LeftPad(c rune, s string, sz int) string {
	return RepeatChars(string(c), sz-len(s)) + s
}
