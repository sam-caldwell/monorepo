package bitfile

/*
 * Bitfile.ReadBytes() method
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * ReadBytes() will read a block of bytes from a given file
 * and return a Block struct from which individual bits can be read
 * or other block operations can be performed.
 */

import (
	"github.com/sam-caldwell/monorepo/go/fs/file/bitBlock"
)

// ReadBytes - Read a sequence of n bytes from the current file and return them as a Block.
func (o *BitFile) ReadBytes(n uint) (block *bitBlock.Block, err error) {
	block = bitBlock.NewBlock(n)
	data := make([]byte, n)
	_, err = o.file.Read(data)
	return block, err
}
