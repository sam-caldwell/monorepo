package types

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/file"
)

// FilePath - File/path data type
type FilePath string

// String - return the string version of a file/path
func (fp *FilePath) String() string {
	return string(*fp)
}

// FromString - given a string, validate the input as a file/path. If mustExist, then verify file exists.
func (fp *FilePath) FromString(path string) {
	*fp = FilePath(path)
}

// FromStringRequireExists - given a string, validate the input as a file/path. If mustExist, then verify file exists.
func (fp *FilePath) FromStringRequireExists(path string, mustExist bool) (err error) {
	fp.FromString(path)
	if mustExist && file.Exists(path) {
		err = fmt.Errorf("file not found")
	}
	return err
}
