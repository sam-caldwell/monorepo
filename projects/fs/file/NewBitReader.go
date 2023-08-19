package file

// NewBitReader - initialize and return a new BitReader object reference
func NewBitReader(bitBufferSize, readBufferSize int) *BitReader {
	var reader BitReader
	reader.count = 0
	reader.done = false
	reader.buffer = make(chan bool, bitBufferSize)
	reader.readBuffer = make([]byte, readBufferSize)
	return &reader
}
