package crsce

/*
 * Block.RangeSha512() method
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * Given a block of related bytes, this method will
 * return the []byte Sha512 hash of a given range of
 * the block's bytes
 */
import (
	"crypto/sha512"
	"fmt"
)

// RangeSha512 - Calculate and return the Sha512 hash of a range of bytes within a given block of bytes
func (blk *Block) RangeSha512(start, stop int) (hash []byte, err error) {
	if stop > len(blk.buffer) {
		return nil, fmt.Errorf("bounds check error")
	}

	if start < 0 {
		return nil, fmt.Errorf("bounds check error (start)")
	}

	if stop < 0 {
		return nil, fmt.Errorf("bounds check error (stop)")
	}

	if start >= stop {
		return nil, fmt.Errorf("stop exceeds start")
	}

	h := sha512.Sum512(blk.buffer[start:stop])

	return h[:], nil

}
