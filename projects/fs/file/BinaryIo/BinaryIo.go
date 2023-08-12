package file

import "os"

// BinaryIo - A binary file reader.
type BinaryIo struct {
	// handle - File handle for the current file.
	handle *os.File
}
