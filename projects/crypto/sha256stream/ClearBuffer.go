package crypto

/*
 * projects/crypto/ClearBuffer.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides a ClearBuffer() method which will
 * pad our 512-bit block buffer.
 */

// ClearBuffer - Clear the buffer and reset our index variables.
func (hash *Sha256Stream) ClearBuffer(lock bool) *Sha256Stream {
	if lock {
		hash.lock.Lock()
		defer hash.lock.Unlock()
	}
	for i := 0; i < len(hash.buffer); i++ {
		hash.buffer[i] = 0
	}
	hash.bitNdx = 0
	hash.byteNdx = 0
	hash.size += uint32(len(hash.buffer))
	return hash
}
