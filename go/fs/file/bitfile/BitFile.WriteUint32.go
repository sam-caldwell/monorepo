package bitfile

/*
 * CRSCE bitfile uint32 writer method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * write a 32-bit unsigned integer to the bitfile.
 */

import (
	"encoding/binary"
	"unsafe"
)

// WriteUint32 - write a 32-bit unsigned integer to the bitfile.
func (o *BitFile) WriteUint32(i uint32) error {
	// Encode the uint64 into a byte slice
	buf := make([]byte, unsafe.Sizeof(i))
	binary.LittleEndian.PutUint32(buf, i)
	return o.WriteBytes(buf)
}
