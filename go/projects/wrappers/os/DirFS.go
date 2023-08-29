package os

import (
	"os"
)

// DirFS - Abstract os.DirFS
var DirFS = os.DirFS

// ResetOsDirFSWrapper - Reset our os.DirFS wrapper to its original native state
func ResetOsDirFSWrapper() {
	DirFS = os.DirFS
}
