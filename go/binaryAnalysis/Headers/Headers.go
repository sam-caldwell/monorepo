package executables

import (
    "errors"
    "fmt"
    "os"
)

type FileHeader interface {
    Elf32 | Elf64
}

// Headers - Generic struct for storing header information
type Headers struct {
	fileHeader FileHeader
    fh *os.File

	//// GetValue - Method signature to return a header data field value
	//GetValue(v ElfHeaderFieldNames) any
}

func (f *Headers) Load(fh *os.File) error{
    var ident [16]byte
    if fh == nil {
        return fmt.Errorf("nil file handle")
    }
    if _, err := fh.ReadAt(ident[:], 0); err != nil {
        return err
    }
    // Verify ELF magic number
    if ident[0] != 0x7f || ident[1] != 'E' || ident[2] != 'L' || ident[3] != 'F' || ident[4] != 2 {
        f.fileHeader=Elf64{}
    }
    return fmt.Errorf("unsupported file type (%s)",fh.Name())
}

func (f *)
