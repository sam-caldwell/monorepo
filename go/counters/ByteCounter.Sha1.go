package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt

 * ByteCounter will create an array of n bytes
 * and allow the count to be incremented 0...255
 * for each byte, carrying 1 to the n+1 byte.
 */

import (
	"crypto/sha1"
	"encoding/hex"
)

// Sha1 - create byte counter hash
func (c *ByteCounter) Sha1() string {
	hash := sha1.Sum(c.v)
	return hex.EncodeToString(hash[:])
}

func (c *ByteCounter) Sha1Bytes() [20]byte {
	return sha1.Sum(c.v)
}
