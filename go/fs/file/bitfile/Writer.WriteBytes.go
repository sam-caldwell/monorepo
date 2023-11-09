package bitfile

/*
 * CRSCE BitFile WriteBytes() Method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

// WriteBytes - Write an array of bits to the bit buffer and flush to disk when buffer is full
func (o *BitFile) WriteBytes(data []byte) (err error) {

	if data != nil {
		for _, currByte := range data {
			if err := o.WriteBit(0x00, currByte); err != nil {
				break
			}
			if err := o.WriteBit(0x01, currByte); err != nil {
				break
			}
			if err := o.WriteBit(0x02, currByte); err != nil {
				break
			}
			if err := o.WriteBit(0x03, currByte); err != nil {
				break
			}
			if err := o.WriteBit(0x04, currByte); err != nil {
				break
			}
			if err := o.WriteBit(0x05, currByte); err != nil {
				break
			}
			if err := o.WriteBit(0x06, currByte); err != nil {
				break
			}
			if err := o.WriteBit(0x07, currByte); err != nil {
				break
			}
		}
	}

	return err
}
