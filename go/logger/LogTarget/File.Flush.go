package LogTarget

// Flush - Flush the log message and reset ANSI colors
//
//	(c) 2023 Sam Caldwell.  MIT License
func (out File) Flush() {
	out.file.Flush()
}
