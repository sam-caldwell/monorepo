package bitreader

import "io"

// NewBitStreamReader - create a new bit-stream reader
//
//	(c) 2023 Sam Caldwell.  MIT License
func NewBitStreamReader(reader io.Reader) *BitStreamReader {
	return &BitStreamReader{
		reader: reader,
		buffer: 0,
		count:  8,
	}
}
