package file

/*
 * projects/fs/file/Reader.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * The file reader struct.
 */

import "os"

// Reader - A file reader
type Reader struct {

	// h - File handle
	h *os.File

	// buffer - byte array used to buffer reads and reduce disk I/O
	buffer []byte

	// filePosition - byte position in the file indicating the next byte to read.
	filePosition uint64

	// bufferPosition - byte position in the buffer. This is the index of the next byte to return
	bufferPosition int

	// bufferBitPos - index of the next bit to be returned from buffer[bufPos] when GetNextBit() is called.
	bufferBitPos int8
}
