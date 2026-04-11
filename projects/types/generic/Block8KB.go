package generic

import "encoding/hex"

type Block8KB [8192]byte

// FromSlice - Given a block of []byte, store to our internal state
func (block *Block8KB) FromSlice(b []byte) {
	copy(block[:], b)
}

// ToSlice - return []byte representation of the internal state
func (block *Block8KB) ToSlice() []byte {
	b := make([]byte, len(*block))
	copy(b, block[:])
	return b
}

// SizeOf - return the size of the object
// Note that we return an unsigned integer because a negative link would not be rational
func (block *Block8KB) SizeOf() uint {
	return uint(len(*block))
}

// String - return the contents of the hash object as a hex string
func (block *Block8KB) String() string {
	return hex.EncodeToString(block[:])
}
