package crsce

/*
 * projects/crsce/Crsce.OpenSource.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Open the CRSCE source file with a given filename.
 * This will return any error state that occurs or it
 * will leave the internal reader open in the virtual
 * csm object.
 */
const (
	// readBufferSize - The number of bytes to be read per I/O cycle.
	readBufferSize = 16384
)

// OpenSource - Open the uncompressed file and set up the bit stream reader to feed compress or decompress.
func (c *Crsce) OpenSource(fileName string) (err error) {
	_, err = c.csm.buffer.Open(fileName, readBufferSize)
	return err
}
