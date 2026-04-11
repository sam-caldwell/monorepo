package hashes

import (
	"crypto/sha512"
	"encoding/hex"
)

const (
	Sha512Length = 64
)

type Sha512 [Sha512Length]byte

// FromSlice - Given a block of []byte, store to our internal state
func (block *Sha512) FromSlice(b []byte) {
	copy(block[:], b)
}

// ToSlice - return []byte representation of the internal state
func (block *Sha512) ToSlice() []byte {
	b := make([]byte, Sha512Length)
	copy(b, block[:])
	return b
}

// SizeOf - return the size of the object
// Note that we return an unsigned integer because a negative link would not be rational
func (block *Sha512) SizeOf() uint {
	return uint(len(*block))
}

// HashString - hash the input and store as state
func (block *Sha512) HashString(data string) {
	block.HashBytes([]byte(data))
}

// HashBytes - hash the input and store as state
func (block *Sha512) HashBytes(data []byte) {
	*block = sha512.Sum512(data)
}

// String - return the contents of the hash object as a hex string
func (block *Sha512) String() string {
	return hex.EncodeToString(block[:])
}

// HexEncodedString - given a hexadecimal-encoded string, return the hash bytes
func (block *Sha512) HexEncodedString(s *string) {
	hash, err := hex.DecodeString(*s)
	if err != nil {
		panic(err)
	}
	*block = [Sha512Length]byte(hash[:Sha512Length])
}
