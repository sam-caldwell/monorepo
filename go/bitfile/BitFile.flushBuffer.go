package bitfile

/*
 * CRSCE bitfile writer
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

func (o *BitFile) flushBuffer() error {
	if err := o.WriteBytes(o.buffer); err != nil {
		return err
	}
	o.buffer = make([]byte, o.bufferSize)
	o.bufferPos = 0
	return nil
}
