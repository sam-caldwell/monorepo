package bitfile

// setBit - set the bit at bitPos in the byte at bufferPos
func (o *Writer) setBit() {
	var writeMask = byte(1 << o.bitPos)
	o.buffer[o.bufferPos] = o.buffer[o.bufferPos] | writeMask
}
