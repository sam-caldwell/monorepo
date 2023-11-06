package crsce

import "github.com/sam-caldwell/monorepo/go/bitfile"

// Compress - given a block of bytes, compress the block and write it to target.
func Compress(block []byte, target bitfile.BitFile) error {
	LSM := make([]byte, s)
	VSM := make([]byte, s)

	exit.CheckError(target.WriteInt(blockId))
	exit.CheckError(target.WriteBytes(LSM))
	exit.CheckError(target.WriteBytes(VSM))
	return nil
}