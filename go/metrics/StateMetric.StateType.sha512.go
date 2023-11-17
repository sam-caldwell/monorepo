package metrics

type Sha512Hash [128]byte

// FromSlice - Given a block of []byte, store to our internal state
func (block *Sha512Hash) FromSlice(b []byte) {
	copy(block[:], b)
}

// ToSlice - return []byte representation of the internal state
func (block *Sha512Hash) ToSlice() (b []byte) {
	copy(b, block[:])
	return b
}

// SizeOf - return the size of the object
// Note that we return an unsigned integer because a negative link would not be rational
func (block *Sha512Hash) SizeOf() uint {
	return uint(len(*block))
}
