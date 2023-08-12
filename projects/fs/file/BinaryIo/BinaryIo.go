package file

import "os"

// BinaryIo - A binary file reader.
type BinaryIo struct {
	// handle - File handle for the current file.
	handle *os.File

	// Buffer read from file
	buffer []byte

	// bufferSize - The number of bytes to read/write from/to buffer
	bufferSize int

	// filePosition - byte position in file marking start of buffer.
	//                buffer[0] is at the current filePosition.
	filePosition uint64

	// hasChanges - counts the number of changes in the buffer that
	//              have not been flushed to disk.
	hasChanges int
}
