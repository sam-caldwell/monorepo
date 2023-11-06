package bitfile

/*
 * CRSCE ReadBytes
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * ReadBytes() will read a block of bytes from a given file
 * and return a Block struct from which individual bits can be read
 * or other block operations can be performed.
 */

// ReadBytes - Read a sequence of n bytes from the current file and return them as a Block.
func (o *BitFile) ReadBytes(n uint) (blk Block, err error) {
	blk.buffer = make([]byte, n)
	_, err = o.file.Read(blk.buffer)
	return blk, err
}
