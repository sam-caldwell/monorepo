package bitwriter

// WriteBit - write a bit to the bit stream
//
//	(c) 2023 Sam Caldwell.  MIT License
func (bsw *BitStreamWriter) WriteBit(bit bool) error {
	if bit {
		bsw.buffer |= 1 << (7 - bsw.count)
	}
	bsw.count++
	if bsw.count == 8 {
		if _, err := bsw.writer.Write([]byte{bsw.buffer}); err != nil {
			return err
		}
		bsw.buffer = 0
		bsw.count = 0
	}
	return nil
}
