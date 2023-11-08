package crsce

import (
	"github.com/sam-caldwell/monorepo/go/fs/file"
)

// writeFileHeader - Write the CRSCE file header.
func writeFileHeader(target *file.BitFile, fileSize int64,
	blockSize uint) error {
	if err := target.WriteInt64(fileSize); err != nil {
		return err
	}
	if err := target.WriteUint(blockSize); err != nil {
		return err
	}
	if err := target.WriteInt(s); err != nil {
		return err
	}
	return nil
}
