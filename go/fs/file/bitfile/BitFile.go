package bitfile

/*
 * CRSCE bitfile reader
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

	// byte position within the file
	filePos int64 // byte position in file

	// bit position in the current buffer element (byte)
	bitPos int

	// buffer containing bytes read into the BitFile periodically.
	buffer []byte

	// byte position in buffer
	bufferPos int

	// actual byte size of the buffer (what was read during last BitFile.Read() operation)
	bufferSize int
}
