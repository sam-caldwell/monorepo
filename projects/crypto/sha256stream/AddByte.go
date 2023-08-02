package crypto

/*
 * projects/crypto/AddBit.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file adds a bit to the buffer and if the buffer is
 * full, the buffer (message block) will be processed into
 * the hash state.
 */

// AddByte - Add a byte to our buffer.
func (hash *Sha256Stream) AddByte(b byte) *Sha256Stream {

	hash.lock.Lock()
	defer hash.lock.Unlock()

	if hash.bitNdx != 0 {
		//
		// We are at a non-zero bit offset
		// So we will pull our byte into the buffer bit for bit
		// since it is not aligned.  Hint: don't use AddByte()
		// if you are not aligned as it *may* have a performance
		// penalty.
		//
		for i := 0; i < 8; i++ {
			bit := b & (1 << i)
			hash.AddBit(bit == 1)
		}
		return hash
	}
	//
	// We are at bitNdx 0 so we are byte-aligned.
	//
	hash.buffer[hash.byteNdx] = b
	hash.byteNdx++

	if hash.byteNdx >= int8(len(hash.buffer)) {
		hash.processMessageBlock()
		hash.ClearBuffer(false)
	}

	return hash

}
