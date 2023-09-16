package ansi

import (
	"fmt"
)

// Debugf - Print debug line (with cobra integration for --debug flag)
func Debugf(format string, msg ...interface{}) {
	Debug(fmt.Sprintf(format, msg...))
}
