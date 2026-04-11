package executables

import (
    "fmt"
    "os"
)

// FileHeader
//
//	(c) 2023 Sam Caldwell.  MIT License
type FileHeader interface {
	Elf32 | Elf64
}

// Headers - Generic struct for storing header information
//
//	(c) 2023 Sam Caldwell.  MIT License
type Headers struct {
	fileHeader FileHeader
	fh         *os.File

	//// GetValue - Method signature to return a header data field value
	//GetValue(v ElfHeaderFieldNames) any
}

// Load
//
//	(c) 2023 Sam Caldwell.  MIT License
func (f *Headers) Load(fh *os.File) error {
	var ident [16]byte
	if fh == nil {
		return fmt.Errorf("nil file handle")
	}
	if _, err := fh.ReadAt(ident[:], 0); err != nil {
		return err
	}
	// Verify ELF magic number
	if ident[0] != 0x7f || ident[1] != 'E' || ident[2] != 'L' || ident[3] != 'F' || ident[4] != 2 {
		f.fileHeader = Elf64{}
	}
	return fmt.Errorf("unsupported file type (%s)", fh.Name())
}
