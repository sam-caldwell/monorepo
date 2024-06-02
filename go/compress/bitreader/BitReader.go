package bitreader

import (
	"io"
)

// BitReader - allow bit-by-bit reading from an io.Reader.
//
//	 This is a mechanism used for bitwise operations against a data source.
//
//		(c) 2023 Sam Caldwell.  MIT License
type BitReader struct {
	//reader - Data source
	reader io.ByteReader
	//buffer - internal buffer
	buffer uint64
	//number of bits in buffer
	numberBitsInBuffer uint64
	//error state
	err error
}
