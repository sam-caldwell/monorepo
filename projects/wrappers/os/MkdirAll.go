package os

import (
	"os"
)

// MkdirAll - Abstract os.MkdirAll
var MkdirAll = os.MkdirAll

// ResetOsMkdirAllWrapper - Reset our os.MkdirAll wrapper to its original native state
func ResetOsMkdirAllWrapper() {
	MkdirAll = os.MkdirAll
}
