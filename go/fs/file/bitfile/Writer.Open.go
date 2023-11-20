package bitfile

/*
 * Writer.Create() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create a file to target for write operations.
 */

import (
	"os"
)

// Open - Open/create a file.  If the file exists, truncate it and open for write.
// The file mode (perms) will be set to 0666.
func (o *Writer) Open(fileName *string, bufferSize int) (err error) {

	// Make sure we don't do anything insane.
	if fileName == nil {
		panic("nil filename encountered by bitfile Writer")
	}

	// Revert to the defaultBufferSize
	if bufferSize == 0 {
		bufferSize = defaultBufferSize
	}

	// Reset all internal state
	o.bufferPos = 0
	o.bitPos = 0

	// Create the buffer
	o.buffer = make([]byte, bufferSize)

	// Create/truncate the buffer file.
	o.file, err = os.Create(*fileName)
	return err
}
