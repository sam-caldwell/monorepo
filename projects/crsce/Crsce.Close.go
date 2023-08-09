package crsce

// Close - Close any file handles
func (crsce *Crsce) Close() {
	defer crsce.csm.buffer.Close()
}
