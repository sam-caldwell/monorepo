package bitfile

import "fmt"

/*
 * Block.Set() method
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * The block class has a Set() method to load data
 * into the buffer.
 */

func (blk *Block) Set(b []byte) (err error) {
	if b == nil {
		return fmt.Errorf("invalid nil input")
	}
	if blk.buffer == nil {
		blk.buffer = make([]byte, len(b))
	}
	copy(blk.buffer, b)
	return nil
}
