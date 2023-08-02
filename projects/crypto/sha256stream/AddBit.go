package crypto

/*
 * projects/crypto/AddBit.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file adds a bit to the buffer and if the buffer is
 * full, the buffer (message block) will be processed into
 * the hash state.
 */

func (hash *Sha256Stream) AddBit(b bool) *Sha256Stream {
	hash.lock.Lock()
	defer hash.lock.Unlock()

	hash.bitNdx++

	if hash.bitNdx == 8 && hash.byteNdx == int8(len(hash.buffer)) {
		hash.processMessageBlock()
		hash.ClearBuffer(false)
	}

	return hash

}
