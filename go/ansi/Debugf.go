package ansi

import (
	"fmt"
)

// Debugf - Print debug line (with cobra integration for --debug flag)
//
//	(c) 2023 Sam Caldwell.  MIT License
func Debugf(format string, msg ...interface{}) {
	Debug(fmt.Sprintf(format, msg...))
}
