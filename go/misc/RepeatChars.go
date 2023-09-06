package misc

// RepeatChars - return a string of s (string) repeated num (int) times.
func RepeatChars(s string, num int) (result string) {
	for i := 0; i < num; i++ {
		result += s
	}
	return result
}
