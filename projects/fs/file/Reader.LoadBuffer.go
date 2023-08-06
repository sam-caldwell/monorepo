package file

import (
	"fmt"
)

// LoadBuffer - Load the Reader.buffer byte array from the file (f.h) at f.fPos if empty
func (f *Reader) LoadBuffer(bufferSize int) error {
	var err error
	var bytesRead int

	// If the buffer is not allocated, allocate it.  Make sure we don't have a null pointer
	if (f.buffer != nil) && (f.bufferPosition < len(f.buffer)) {
		return nil
	}
	// We have an empty buffer, load content.

	// Reset buffer index and reallocate the buffer
	f.bufferPosition = 0
	f.buffer = make([]byte, bufferSize)

	// Read from the file starting from f.fPos into the buffer
	if bytesRead, err = f.h.ReadAt(f.buffer, int64(f.filePosition)); err != nil {
		return err
	}

	// Update positions after successful read
	fmt.Printf("bytesRead: %d\n", bytesRead)
	f.filePosition += uint64(bytesRead)
	f.bufferPosition = 0
	f.bufferBitPos = 0
	fmt.Printf("pos:%d\n", f.filePosition)
	return nil
}
