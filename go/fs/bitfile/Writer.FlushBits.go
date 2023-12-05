package bitfile

// FlushBits - Flush a set of bits as a complete byte and increment the buffer
func (o *Writer) FlushBits() error {
	o.bufferPos++
	if int(o.bufferPos) >= len(o.buffer) {
		return o.FlushBytes()
	}
	o.bitPos = 0
	return nil
}
