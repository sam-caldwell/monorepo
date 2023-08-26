package file

/*
 * projects/fs/file/GetFileSize.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return a file's size using the file handle
 */

import (
	"github.com/sam-caldwell/go/v2/projects/go/ansi"
	"os"
)

// GetFileSize - Return the int64 file size of a given file.
func GetFileSize(f *os.File) int64 {
	fileInfo, err := f.Stat()
	if err != nil {
		ansi.Red().Println(err.Error()).Fatal(1)
	}
	return fileInfo.Size()
}
