package executables

import "os"

// Load - ELF 64-bit header load method.
func (h *Elf64) Load(fh *os.File) (err error) {
	//Todo:
	//  - Load the file header
	//  - Parse the data into the appropriate fields
	//  - Return any error state
	var originalPosition int64
	originalPosition, err = fh.Seek(0, 0)
	if err != nil {
		return err
	}

	defer func() {
		_, err = fh.Seek(originalPosition, 0)
	}()
	return err
}
