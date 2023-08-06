package file

// Close - Close the file handle if it exists and is open
func (f *Reader) Close() {
	if f.h != nil {
		_ = f.h.Close()
		f.h = nil
	}
}
