package metrics

type Block1KB [1024]byte

// FromSlice - Given a block of []byte, store to our internal state
func (block *Block1KB) FromSlice(b []byte) {
	copy(block[:], b)
}

// ToSlice - return []byte representation of the internal state
func (block *Block1KB) ToSlice() (b []byte) {
	copy(b, block[:])
	return b
}

// SizeOf - return the size of the object
// Note that we return an unsigned integer because a negative link would not be rational
func (block *Block1KB) SizeOf() uint {
	return uint(len(*block))
}
