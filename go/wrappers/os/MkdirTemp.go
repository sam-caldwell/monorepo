package os

import (
	"os"
)

// MkdirTemp - Abstract os.MkdirTemp
var MkdirTemp = os.MkdirTemp

// ResetOsMkdirTempWrapper - Reset our os.MkdirTemp wrapper to its original native state
func ResetOsMkdirTempWrapper() {
	MkdirTemp = os.MkdirTemp
}
