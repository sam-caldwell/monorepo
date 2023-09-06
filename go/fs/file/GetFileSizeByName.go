package file

/*
 * projects/fs/file/GetFileSizeByName.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return a file's size using the filename
 */

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os"
)

// GetFileSizeByName - Return the int64 file size of a given file (by filename).
func GetFileSizeByName(f string) int64 {
	handle, err := os.Open(f)
	if err != nil {
		ansi.Red().Println(err.Error()).Fatal(1)
	}
	defer func() { _ = handle.Close() }()
	return GetFileSize(handle)
}
