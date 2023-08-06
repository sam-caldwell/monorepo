package file

const (
	ClearBit = false
	SetBit   = true
)

// GetNextBit - Return next bit from the current byte from buffer (call LoadBuffer() if needed).
func (f *Reader) GetNextBit() (out bool, err error) {
	// Check if the buffer is empty or all bytes in the buffer have been consumed,
	// or if we have consumed all bits in the current byte.
	if f.bufferPosition >= len(f.buffer) {
		//	Call LoadBuffer() to fill the buffer from the file.
		if err := f.LoadBuffer(len(f.buffer)); err != nil {
			return ClearBit, err
		}
	}
	// Check if there are bytes left in the buffer to read bits from.
	if f.bufferPosition < len(f.buffer) {
		//Get the current byte from the buffer.
		currentByte := f.buffer[f.bufferPosition]
		// Get the next bit from the current byte and update bitPos accordingly.
		out = (currentByte>>uint(7-f.bufferBitPos))&1 != 0x00
		f.bufferBitPos++
		// If we have consumed all bits in the current byte, move to the next byte.
		if f.bufferBitPos >= 8 {
			f.bufferPosition++
			f.bufferBitPos = 0
		}
	}
	return out, nil
}
