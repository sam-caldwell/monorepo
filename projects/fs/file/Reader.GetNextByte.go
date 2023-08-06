package file

// GetNextByte - Return next byte from buffer (call LoadBuffer() if needed).
func (f *Reader) GetNextByte() (out byte, err error) {
	// Check if buffer is empty or all bytes in the buffer have been consumed.
	//if f.bufferPosition >= len(f.buffer) {
	// Call LoadBuffer() to fill the buffer from the file.
	//if err := f.LoadBuffer(); err != nil {
	//	return 0, err
	//}
	//}
	// Check if there are bytes left in the buffer to return.
	//if f.bufferPosition < len(f.buffer) {
	// Get the next byte from the buffer.
	//out = f.buffer[f.bufferPosition]
	// Move the buffer position to point to the next byte.
	//f.bufferPosition++
	//}
	return out, nil
}
