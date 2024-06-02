package bitreader

// ReadBits - Read a number of bits from the BitReader and return as an unsigned integer.
//
// This is an alias of ReadBits64().
//
//	(c) 2023 Sam Caldwell.  MIT License
func (br *BitReader) ReadBits(requestedNumberOfBits uint8) (n uint64) {
	return br.ReadBits64(requestedNumberOfBits)
}
