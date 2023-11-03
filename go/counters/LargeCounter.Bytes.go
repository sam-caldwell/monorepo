package counters

import "encoding/binary"

// Bytes - Return the bytes in reverse order
func (c *LargeCounter) Bytes() (out []byte) {
	const (
		byteSize     = 8 // 8 bits per byte
		bytesPerWord = 8 // 8 bytes per 64-bit word
	)

	out = make([]byte, byteSize*len(*c))
	for i := 0; i < len(*c); i++ {
		this := make([]byte, bytesPerWord)
		binary.BigEndian.PutUint64(this, (*c)[i])
		out = append(out[:i*byteSize], this...)
	}

	return out
}
