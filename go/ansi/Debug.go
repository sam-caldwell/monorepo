package ansi

import (
	"fmt"
)

// Debug - Print debug message (with cobra integration for --debug flag)
func Debug(msg interface{}) {
	if debugMode {
		Yellow().Print(fmt.Sprintf("[Debug]: %v", msg)).Reset()
	}
}
