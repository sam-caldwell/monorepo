package bitreader

// ReadBit - read one bit at a time from a file.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (bsr *BitStreamReader) ReadBit() (bool, error) {
	if bsr.count == 8 {
		buf := make([]byte, 1)
		if _, err := bsr.reader.Read(buf); err != nil {
			return false, err
		}
		bsr.buffer = buf[0]
		bsr.count = 0
	}
	bit := (bsr.buffer & (1 << (7 - bsr.count))) != 0
	bsr.count++
	return bit, nil
}
