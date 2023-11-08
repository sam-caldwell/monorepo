package bitBlock

/*
 * Block.Set() method
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * The block class has a Set() method to load data
 * into the buffer.
 */

func (block *Block) Set(b []byte) {
	if b == nil {
		b = make([]byte, 0)
	}
	if block.buffer == nil {
		block.buffer = make([]byte, len(b))
	}
	copy(block.buffer, b)
}
