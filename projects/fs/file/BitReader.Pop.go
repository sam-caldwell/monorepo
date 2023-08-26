package file

// Pop - Pop a single bit from the buffer and the done status.
func (reader *BitReader) Pop() (result bool, eof bool) {
	return <-reader.buffer, reader.done
}
