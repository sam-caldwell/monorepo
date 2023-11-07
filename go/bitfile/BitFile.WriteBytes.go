package bitfile

/*
 * CRSCE bitfile writer
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

// WriteBytes - Write an array of bits to the bit buffer and flush to disk when buffer is full
func (o *BitFile) WriteBytes(data []byte) error {
	if data == nil {
		return nil //Do nothing
	}
	for currByte := range data {
		for b := 0; b < 8; b++ {
			mask := 2 << b
			bit := (currByte & mask) != 0
			if err := o.Write(bit); err != nil {
				return err
			}
		}
	}
	return nil
}
