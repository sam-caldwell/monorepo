package bitwriter

import "io"

// NewBitStreamWriter - Create a new bit stream writer
//
//	(c) 2023 Sam Caldwell.  MIT License
func NewBitStreamWriter(writer io.Writer) *BitStreamWriter {
	return &BitStreamWriter{
		writer: writer,
		buffer: 0,
		count:  0,
	}
}
