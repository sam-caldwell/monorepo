package hashes

import (
	"crypto/sha1"
	"encoding/hex"
)

const (
	Sha1Length = 20
)

type Sha1 [Sha1Length]byte

// FromSlice - Given a block of []byte, store to our internal state
func (block *Sha1) FromSlice(b []byte) {
	copy(block[:], b)
}

// ToSlice - return []byte representation of the internal state
func (block *Sha1) ToSlice() []byte {
	b := make([]byte, Sha1Length)
	copy(b, block[:])
	return b
}

// SizeOf - return the size of the object
// Note that we return an unsigned integer because a negative link would not be rational
func (block *Sha1) SizeOf() uint {
	return uint(len(*block))
}

// HashString - hash the input and store as state
func (block *Sha1) HashString(data string) {
	block.HashBytes([]byte(data))
}

// HashBytes - hash the input and store as state
func (block *Sha1) HashBytes(data []byte) {
	*block = sha1.Sum([]byte(data))
}

// String - return the contents of the hash object as a hex string
func (block *Sha1) String() string {
	return hex.EncodeToString(block[:])
}
