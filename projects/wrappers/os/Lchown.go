package os

import (
	"os"
)

// Lchown - Abstract os.Lchown
var Lchown = os.Lchown

// ResetOsLchownWrapper - Reset our os.Lchown wrapper to its original native state
func ResetOsLchownWrapper() {
	Lchown = os.Lchown
}
