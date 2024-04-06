package binaryAnalysis

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os"
)

type FileFormat uint8

const (
	unknown = iota
	Elf32
	Elf64
	linuxSo32
	linuxSo64
	Pe32
	Pe64
	Dll32
	Dll64
)

func GetFileType(fh *os.File, debug bool) (FileFormat, error) {
	var header [20]byte

	if _, err := fh.ReadAt(header[:], 0); err != nil {
		return unknown, err
	}
	if header[0] == 0x7f || header[1] == 'E' || header[2] == 'L' || header[3] == 'F' || header[4] == 2 {
		ansi.Cyan()
		for pos, value := range header {
			ansi.Printf("[%02x:%02x]", pos, value)
		}
		ansi.LF().Reset()
		return Elf64, nil
	}
	return unknown, fmt.Errorf("unknown or unsupported file format")
}
