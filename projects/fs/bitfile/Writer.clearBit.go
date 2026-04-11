package bitfile

// clearBit - clear the bit at bitPos in the byte at bufferPos
func (o *Writer) clearBit() {
	var writeMask = ^byte(1 << o.bitPos)
	o.buffer[o.bufferPos] = o.buffer[o.bufferPos] & writeMask
}
