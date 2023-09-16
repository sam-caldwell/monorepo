package ansi

// Debugln - Print debug line with LF ending (with cobra integration for --debug flag)
func Debugln(format string, msg ...interface{}) {
	Debug(msg)
	LF()
}
