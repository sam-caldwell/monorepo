package bitfile

/*
 * bitfile Writer
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit writer
 */

import (
	"os"
)

// Writer - A Bitfile struct represents a single bitwise writer object
type Writer struct {

	// file handle for the associated file
	file *os.File

	// A byte buffer used to assemble a number of bytes
	buffer []byte

	// The Buffer position tracks the current byte in the buffer
	bufferPos uint16

	// The bit position used to write bits into a byte before it is pushed to the buffer
	bitPos byte
}
