package executables

import (
	"encoding/binary"
	"errors"
	"os"
)

// Load loads Elf64 header from the given file handle.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (elf *Elf64) Load(fh *os.File) error {
	// Read the ELF identification field to verify it's an ELF file
	var ident [16]byte
	if _, err := fh.ReadAt(ident[:], 0); err != nil {
		return err
	}
	// Verify ELF magic number
	if ident[0] != 0x7f || ident[1] != 'E' || ident[2] != 'L' || ident[3] != 'F' || ident[4] != 2 {
		return errors.New("not an ELF 64-bit binary")
	}

	// Reset the file position to the start of the file
	if _, err := fh.Seek(0, 0); err != nil {
		return err
	}

	// Read ELF header
	if err := binary.Read(fh, binary.LittleEndian, elf); err != nil {
		return err
	}

	return nil
}
