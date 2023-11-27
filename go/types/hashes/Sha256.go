package hashes

import (
	"crypto/sha256"
	"encoding/hex"
)

const (
	Sha256Length = 32
)

type Sha256 [Sha256Length]byte

// FromSlice - Given a block of []byte, store to our internal state
func (block *Sha256) FromSlice(b []byte) {
	copy(block[:], b)
}

// ToSlice - return []byte representation of the internal state
func (block *Sha256) ToSlice() []byte {
	b := make([]byte, Sha256Length)
	copy(b, block[:])
	return b
}

// SizeOf - return the size of the object
// Note that we return an unsigned integer because a negative link would not be rational
func (block *Sha256) SizeOf() uint {
	return uint(len(*block))
}

// HashString - hash the input and store as state
func (block *Sha256) HashString(data string) {
	block.HashBytes([]byte(data))
}

// HashBytes - hash the input and store as state
func (block *Sha256) HashBytes(data []byte) {
	*block = sha256.Sum256(data)
}

// String - return the contents of the hash object as a hex string
func (block *Sha256) String() string {
	return hex.EncodeToString(block[:])
}

// HexEncodedString - given a hexadecimal-encoded string, return the hash bytes
func (block *Sha256) HexEncodedString(s *string) {
	hash, err := hex.DecodeString(*s)
	if err != nil {
		panic(err)
	}
	*block = [Sha256Length]byte(hash[:Sha256Length])
}
