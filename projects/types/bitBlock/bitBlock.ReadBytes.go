package bitBlock

/*
 * bitBlock.ReadBytes() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Return a given number of bytes from a block
 */

// ReadBytes - Return sz number of bytes from the block buffer.  If sz==0, return it all.
func (block *Block) ReadBytes(sz uint) (data []byte) {
	if block.buffer == nil {
		return nil
	}
	if sz == 0 {
		return block.buffer[:]
	}
	return block.buffer[0:sz]
}
