package bitfile

/*
 * bitfile Reader
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader
 */

import (
	"os"
)

// Reader - A Bitfile struct represents a single bitwise reader object
type Reader struct {
	blockSize uint
	// file handle for the associated file
	file *os.File
}
