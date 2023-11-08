package bitfile

/*
 * BitFile.Read() method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Read a single block of data from the file.
 */

import "github.com/sam-caldwell/monorepo/go/fs/file/bitBlock"

// Read - Read a block of bytes from the file and return the block to the caller.
func (o *BitFile) Read() (block *bitBlock.Block, err error) {
	var bytesRead int
	buffer := make([]byte, block.Size())
	bytesRead, err = o.file.Read(buffer)
	if bytesRead > 0 {
		block.Set(buffer[0:bytesRead])
	}
	return block, err
}
