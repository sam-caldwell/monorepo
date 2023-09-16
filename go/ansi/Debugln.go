package ansi

import "fmt"

// Debugln - Print debug line with LF ending (with cobra integration for --debug flag)
func Debugln(msg interface{}) {
	Debug(fmt.Sprintf("%v\n", msg))
}
