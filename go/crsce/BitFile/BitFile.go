package bitfile

/*
 * CRSCE BitFile reader
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

	filePos int64 // byte position in file

	bitPos byte // bit position in current byte (in buffer)

	bufferPos int // byte position in buffer

	bufferSize int // actual byte size of the buffer (what was read)
}
