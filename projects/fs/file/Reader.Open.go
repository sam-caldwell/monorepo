package file

import (
	"fmt"
	"os"
)

// Open - open a file for the current reader and load the buffer
func (f *Reader) Open(name string, bufferSize int) (out *Reader, err error) {

	// Make sure any previously opened file is closed.
	f.Close()

	if !Existsp(&name) {
		return f, fmt.Errorf("file not found")
	}

	// Open the file
	f.h, err = os.Open(name)
	if err != nil {
		return f, err
	}

	// Load the buffer with file contents
	err = f.LoadBuffer(bufferSize)

	// Return file handle and any error
	return f, err
}
