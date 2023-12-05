package bitfile

/*
 * bitfile Writer.WriteUint64() method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * write a 64-bit unsigned integer to the bitfile.
 */

import (
	"encoding/binary"
	"unsafe"
)

// WriteUint64 - write a 64-bit unsigned integer to the bitfile.
func (o *Writer) WriteUint64(i uint64) error {

	// Encode the uint64 into a byte slice

	buf := make([]byte, unsafe.Sizeof(i))
	binary.LittleEndian.PutUint64(buf, i)

	// Write the byte slice to the file
	return o.WriteBytes(buf)
}
