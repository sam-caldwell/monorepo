package bitfile

/*
 * bitfile reader
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

import (
	"os"
)

// BitFile - A Bitfile struct represents a single bitwise reader/writer object
type BitFile struct {
	// file handle for the associated file
	file *os.File
}
