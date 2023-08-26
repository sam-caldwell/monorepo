package file

// Close - Close the bit reader buffer and set done to true
func (reader *BitReader) Close() {
	reader.done = true
	close(reader.buffer)
}
