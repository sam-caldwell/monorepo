package bitreader

// ReadBit - read a single bit from the BitReader data source and return as a boolean.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (br *BitReader) ReadBit() bool {
	bit := br.ReadBits64(1)
	return bit != 0
}
