package bitwriter

// Flush - Flush the bit stream writer
//
//	(c) 2023 Sam Caldwell.  MIT License
func (bsw *BitStreamWriter) Flush() error {
	if bsw.count > 0 {
		if _, err := bsw.writer.Write([]byte{bsw.buffer}); err != nil {
			return err
		}
	}
	return nil
}
