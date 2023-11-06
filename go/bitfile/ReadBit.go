package bitfile

/*
 * Block.ReadBit() method to read block bit by bit.
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Given a block of data, read and return each bit
 * of content.
 */

import (
	"fmt"
	"log"
)

// ReadBit - Read the block one bit at a time, returning the bit at position p and return error if
// position p is out of bounds.
func (blk *Block) ReadBit(p int) (bit bool, err error) {
	if p >= 8*len(blk.buffer) {
		return false, fmt.Errorf("index out of range")
	}
	//Get the byte index
	i := p / 8

	//Get the bit position in byte i
	b := p % 8

	//get this byte
	thisByte := blk.buffer[i]

	mask := byte(1 << b)

	//bit value
	v := thisByte & mask

	log.Printf("pos: %02d  bit: %02d mask: %08b v: %08b", p, b, mask, v)

	//if bit b is not zero, then return true indicating set bit.
	return v != 0, err
}
