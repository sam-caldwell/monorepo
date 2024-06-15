package bitreader

import "io"

// BitStreamReader - allow bit-by-bit reading from an io.Reader.
//
//	 This is a mechanism used for bitwise operations against a data source.
//
//		(c) 2023 Sam Caldwell.  MIT License
type BitStreamReader struct {
	reader io.Reader
	buffer byte
	count  int
}
