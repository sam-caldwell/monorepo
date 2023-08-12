package crsce

// GetNext - Get the next bit in our input data stream.
func (csm *CSM) GetNext() bool {
	out, err := csm.buffer.GetNextBit()
	csm.position++
	csm.err = err
	return out
}
