package os

import (
	"os"
)

// Stdout - Abstract os.Stdout
var Stdout = os.Stdout

// ResetOsStdoutWrapper - Reset our os.Stdout wrapper to its original native state
func ResetOsStdoutWrapper() {
	Stdout = os.Stdout
}
