package bitreader

// Error - Return the current error state
//
//	(c) 2023 Sam Caldwell.  MIT License
func (br *BitReader) Error() error {
	return br.err
}
