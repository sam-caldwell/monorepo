package file

/*
 * CRSCE bitfile reader
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

import (
	"os"
)

const (
	bufferSize = 4096
)

type BitFile struct {
	file *os.File

	buffer []byte

	bitPos int

	filePos int64 // byte position in file

	bufferPos int // byte position in buffer

	bufferSize int // actual byte size of the buffer (what was read)
}
