package counters

import "encoding/binary"

func (c *LargeCounter) Bytes() (out []byte) {
	const (
		bytesPerWord = 8 // 8 bytes per 64-bit word
	)

	sz := len(*c)
	out = make([]byte, sz*bytesPerWord)
	this := make([]byte, bytesPerWord)
	for i := sz - 1; i >= 0; i-- {
		binary.BigEndian.PutUint64(this, (*c)[i])
		copy(out[(sz-1-i)*bytesPerWord:], this)
	}
	return out
}

func (c *LargeCounter) FastBytes() (out []byte) {
	const (
		bytesPerWord = 8 // 8 bytes per 64-bit word
	)

	sz := len(*c) * bytesPerWord
	out = make([]byte, sz*bytesPerWord)
	this := make([]byte, bytesPerWord)
	for i := 0; i < sz; i++ {
		binary.BigEndian.PutUint64(this, (*c)[i])
		copy(out[i*bytesPerWord:], this)
	}
	return out
}
