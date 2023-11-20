package bitBlock

/*
 * bitBlock.RangeSha1() method
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * Given a block of related bytes, this method will
 * return the []byte sha1 hash of a given range of
 * the block's bytes
 */
import (
	"crypto/sha1"
	"fmt"
)

// RangeSha1 - Calculate and return the SHA1 hash of a range of bytes within a given block of bytes
func (block *Block) RangeSha1(start, stop int) (hash []byte, err error) {
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

	h := sha1.Sum(block.buffer[start:stop])

	return h[:], nil

}
