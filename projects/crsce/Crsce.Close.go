package crsce

// Close - Close any file handles
func (c *Crsce) Close() {
	defer c.csm.buffer.Close()
}
