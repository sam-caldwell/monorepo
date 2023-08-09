package crsce

// NewFileCompressor - Create a CRSCE file compressor
func NewFileCompressor(source string, bufferSize int) (out *Crsce, err error) {
	var crsce Crsce
	if _, err = crsce.csm.buffer.Open(source, bufferSize); err != nil {
		return nil, err
	}
	return &crsce, nil
}
