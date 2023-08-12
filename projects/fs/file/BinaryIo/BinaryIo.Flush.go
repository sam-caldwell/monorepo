package file

func (io *BinaryIo) Flush() {
	if io.handle == nil {
		return
	}
	_ = io.handle.Sync()
}
