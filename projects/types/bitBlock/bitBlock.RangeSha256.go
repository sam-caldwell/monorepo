package bitBlock

/*
 * Block.RangeSha256() method
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * Given a block of related bytes, this method will
 * return the []byte Sha256 hash of a given range of
 * the block's bytes
 */
import (
	"crypto/sha256"
	"fmt"
)

// RangeSha256 - Calculate and return the Sha256 hash of a range of bytes within a given block of bytes
func (block *Block) RangeSha256(start, stop int) (hash []byte, err error) {
	block.lock.Lock()
	defer block.lock.Unlock()

	if stop > len(block.buffer) {
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

	h := sha256.Sum256(block.buffer[start:stop])

	return h[:], nil

}
