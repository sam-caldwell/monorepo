package crsce

// Open - Open a file.Reader object to read our source file
func (csm *CSM) Open(fileName string, readBufferSize int) (err error) {
	_, err = csm.buffer.Open(fileName, readBufferSize)
	return err
}
