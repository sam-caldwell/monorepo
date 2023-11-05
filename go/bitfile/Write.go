package bitfile

/*
 * CRSCE bitfile writer
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

// Write - Write file bits to a target file in 4K chunks.
func (o *BitFile) Write() (err error) {
	return nil
}

func (o *BitFile) WriteUint(i uint) error {
	return nil
}

func (o *BitFile) WriteInt(i int) error {
	return nil
}

func (o *BitFile) WriteInt64(i int64) error {
	return nil
}

func (o *BitFile) WriteBytes(b []byte) error {
	if b == nil {
		return nil //Do nothing
	}

	//ToDo: write bytes to the buffer for write to file

	return nil
}
