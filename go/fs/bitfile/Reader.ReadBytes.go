package bitfile

/*
 * Bitfile Reader.ReadBytes() method
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * Reader.ReadBytes() will read a block of bytes from a given file
 * and return a Block struct from which individual bits can be read
 * or other block operations can be performed.
 */

import (
	bitBlock2 "github.com/sam-caldwell/monorepo/go/types/bitBlock"
)

// ReadBytes - Read a sequence of n bytes from the current file and return them as a Block.
func (o *Reader) ReadBytes(n uint) (block *bitBlock2.Block, err error) {

	block = bitBlock2.NewBlock(n)

	data := make([]byte, n)

	_, err = o.file.Read(data)

	block.Set(data)

	return block, err

}
