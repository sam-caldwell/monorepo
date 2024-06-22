package random

import (
	"crypto/rand"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
)

// GenerateRandomFile - Generate a random file of a given size
func GenerateRandomFile(fileName *string, fileSz uint64) (err error) {
	var f *os.File
	if file.Exists(*fileName) {
		_ = file.Delete(*fileName)
	}
	if f, err = os.Create(*fileName); err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	buffer := make([]byte, 4096)

	for written := uint64(0); written < fileSz; {
		var n int
		var toWrite uint64
		if toWrite = fileSz - written; toWrite > uint64(len(buffer)) {
			toWrite = uint64(len(buffer))
		}
		if n, err = rand.Read(buffer[:toWrite]); err != nil {
			return err
		}
		if n == 0 {
			return nil
		}
		if n, err = f.Write(buffer[:n]); err != nil {
			return err
		}
		written += uint64(n)
	}
	return err
}
