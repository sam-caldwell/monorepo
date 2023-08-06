package file

// NewFileReader - initialize and create a new Reader
func NewFileReader(name string, bufferSize int) (o *Reader, err error) {
	return o.Open(name, bufferSize)
}
