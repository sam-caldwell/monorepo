package os

import (
	"os"
)

// Mkdir - Abstract os.Mkdir
var Mkdir = os.Mkdir

// ResetOsMkdirWrapper - Reset our os.Mkdir wrapper to its original native state
func ResetOsMkdirWrapper() {
	Mkdir = os.Mkdir
}
