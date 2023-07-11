package os

import (
	"os"
)

// Exit - Abstract os.Exit
var Exit = ResetOsExitWrapper()

// Stdout - Abstract os.Stdout
var Stdout = ResetOsStdoutWrapper()

// ResetOsExitWrapper - Reset our os.Exit wrapper to its original native state
func ResetOsExitWrapper() func(n int) {
	Exit = os.Exit
	return Exit
}

// ResetOsStdoutWrapper - Reset our os.Stdout wrapper to its original native state
func ResetOsStdoutWrapper() *os.File {
	Stdout = os.Stdout
	return Stdout
}
