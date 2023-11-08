package bitfile

/*
 * CRSCE BitFile.flushBuffer() method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * If the BitFile buffer is full, flush the byte buffer array
 * to disk.
 */

// flushBuffer - flush the full buffer to disk
func (o *BitFile) flushBuffer() error {

	if o.bufferPos == len(o.buffer) {
		defer func() { _ = o.file.Sync() }()
		bytesWritten, err := o.file.Write(o.buffer)
		o.filePos += int64(bytesWritten)
		return err
	}

	o.bufferPos = 0
	o.buffer = make([]byte, o.bufferSize)

	return nil

}
