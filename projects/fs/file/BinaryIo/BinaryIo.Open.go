package file

import (
	"os"
)

// Open - Open the file handle
func (io *BinaryIo) Open(fileName string) (err error) {
	if io.handle != nil {
		_ = io.handle.Close()
	}
	io.handle, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	return err
}
