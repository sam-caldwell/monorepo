package os

import (
	"io/fs"
	"os"
)

// DirFS - Abstract os.DirFS
var DirFS = ResetOsDirFSWrapper()

// ResetOsDirFSWrapper - Reset our os.DirFS wrapper to its original native state
func ResetOsDirFSWrapper() func(dir string) fs.FS {
	DirFS = os.DirFS
	return DirFS
}
