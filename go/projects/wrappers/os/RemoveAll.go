package os

import (
	"os"
)

// RemoveAll - Abstract os.RemoveAll
var RemoveAll = os.RemoveAll

// ResetOsRemoveAllWrapper - Reset our os.RemoveAll wrapper to its original native state
func ResetOsRemoveAllWrapper() {
	RemoveAll = os.RemoveAll
}
