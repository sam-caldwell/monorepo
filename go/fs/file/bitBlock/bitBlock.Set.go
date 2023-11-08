package bitBlock

import "fmt"

/*
 * Block.Set() method
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * The block class has a Set() method to load data
 * into the buffer.
 */

func (block *Block) Set(b []byte) (err error) {
	if b == nil {
		return fmt.Errorf("invalid nil input")
	}
	if block.buffer == nil {
		block.buffer = make([]byte, len(b))
	}
	copy(block.buffer, b)
	return nil
}
