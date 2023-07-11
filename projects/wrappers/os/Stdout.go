package os

import (
	"os"
)

// Stdout - Abstract os.Stdout
var Stdout = ResetOsStdoutWrapper()

// ResetOsStdoutWrapper - Reset our os.Stdout wrapper to its original native state
func ResetOsStdoutWrapper() *os.File {
	Stdout = os.Stdout
	return Stdout
}
