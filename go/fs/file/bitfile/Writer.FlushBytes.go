package bitfile

func (o *Writer) FlushBytes() error {
	if (o.buffer == nil) || (o.file == nil) {
		return nil //Do nothing on nil/empty buffer or non-open file
	}
	_, err := o.file.Write(o.buffer)
	o.bufferPos = 0
	return err
}
