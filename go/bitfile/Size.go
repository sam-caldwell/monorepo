package bitfile

/*
 * CRSCE bitfile size method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return the size of the file
 */

import (
	"fmt"
	"os"
)

// Size - Return file size (in bytes)
func (o *BitFile) Size() (sz int64, err error) {
	if o.file == nil {
		return 0, fmt.Errorf("file unknown")
	}
	fileInfo, err := os.Stat(o.file.Name())
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}
