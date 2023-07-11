package os

import (
	"os"
)

// Pipe - Abstract os.Pipe
var Pipe = os.Pipe

// ResetOsPipeWrapper - Reset our os.Pipe wrapper to its original native state
func ResetOsPipeWrapper() {
	Pipe = os.Pipe
}
