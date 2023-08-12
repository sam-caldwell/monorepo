package file

func (io *BinaryIo) Close() {
	if io.handle == nil {
		return
	}
	_ = io.handle.Sync()
	_ = io.handle.Close()
	io.handle = nil
}
