package bitfile

/*
 * Reader.Read() method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Read a single block of data from the file.
 */

import (
	bitBlock2 "github.com/sam-caldwell/monorepo/go/types/bitBlock"
)

// Read - Read a block of bytes from the file and return the block to the caller.
func (o *Reader) Read() (block *bitBlock2.Block, err error) {

	//var bytesRead int // count of bytes read
	var buffer []byte

	block = bitBlock2.NewBlock(o.blockSize)

	buffer = make([]byte, block.Size())

	if _, err = o.file.Read(buffer); err != nil {
		return nil, err
	}

	block.Set(buffer)

	//if (bytesRead > 0) && (bytesRead < len(buffer)) {
	//	block.Set(buffer)
	//}

	return block, err
}
