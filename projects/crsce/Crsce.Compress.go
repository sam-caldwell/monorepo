package crsce

import "io"

// Compress - Compress the previously opened file and write the resulting byte stream to the given fileName
func (c *Crsce) Compress(fileName string) (err error) {

	c.lsm = make([]CrossSum, getCrossSumSize(c.blockSize))
	c.vsm = make([]CrossSum, getCrossSumSize(c.blockSize))

	var exit bool = false

	for !exit {
		bit := c.csm.GetNext()
		if c.csm.err != nil {
			if c.csm.err != io.EOF {
				return err
			}
			break
		}
		x, y := getCsmCoordinates(c.csm.position, c.blockSize)

		if bit {
			c.lsm[y].Increment()
			c.vsm[x].Increment()
		}

	}
	return nil
}
