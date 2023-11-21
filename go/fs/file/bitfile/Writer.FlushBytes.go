package bitfile

func (o *Writer) FlushBytes() error {
	if (o.buffer == nil) || (o.file == nil) {
		return nil //Do nothing on nil/empty buffer or non-open file
	}
	if _, err := o.file.Write(o.buffer); err != nil {
		return err
	}
	if err := o.file.Sync(); err != nil {
		return err
	}
	o.bufferPos = 0
	return nil
}
