package crypto

/*
 * projects/crypto/AddBit.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file adds a bit to the buffer and if the buffer is
 * full, the buffer (message block) will be processed into
 * the hash state.
 */

// AddBit - Add a bit to our hash stream and flush when we have a byte.
func (hash *Sha256Stream) AddBit(b bool) *Sha256Stream {
	hash.buffer = hash.buffer << 1 //shift left
	if b {
		hash.buffer |= 1
	} else {
		hash.buffer |= 0
	}
	hash.bitNdx++
	if hash.bitNdx == 8 {
		hash.h.Write([]byte{hash.buffer})
		hash.buffer = 0x00
		hash.bitNdx = 0
	}
	return hash
}
