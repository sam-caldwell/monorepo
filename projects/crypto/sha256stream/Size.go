package crypto

/*
 * projects/crypto/Size.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines a struct for the Sha256Stream
 * which allows a user to stream in an arbitrary series
 * of bits or bytes and then compute the hash of the
 * stream.
 */

// Size - Return the number of bytes processed by the hash stream.
func (hash *Sha256Stream) Size() int {
	/*
		This can return the number of bytes processed before Output() is called.
		But it will only return the number of bytes processed (excluding what may)
		be in the buffer.  But the caller can assume that if Output() were called
		at any point the full 512-bit buffer would be flushed into this count.
	*/
	hash.lock.Lock()
	defer hash.lock.Unlock()
	return int(hash.size)
}
