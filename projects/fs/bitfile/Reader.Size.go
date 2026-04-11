package bitfile

/*
 * bitfile Reader.Size() method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return the size of the file
 */

import (
	"fmt"
	"os"
)

// Size - Return file size (in bytes)
func (o *Reader) Size() (sz uint64, err error) {

	if o.file == nil {
		return 0, fmt.Errorf("file unknown")
	}

	fileInfo, err := os.Stat(o.file.Name())

	if err != nil {
		return 0, err
	}

	return uint64(fileInfo.Size()), nil

}
