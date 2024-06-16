package ansi

import "fmt"

// Debugln - Print debug line with LF ending (with cobra integration for --debug flag)
//
//	(c) 2023 Sam Caldwell.  MIT License
func Debugln(msg interface{}) {
	Debug(fmt.Sprintf("%v\n", msg))
}
