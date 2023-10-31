package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * ByteCounter Adder Function
 */

func (c *ByteCounter) Add(value int) (err error) {
	for i := 0; i < value; i++ {
		if c.v[0] == 255 {
			c.v[0] = 0
			if err = c.carry(1); err != nil {
				break
			}
		} else {
			if err = c.inc(0); err != nil {
				break
			}
		}
	}
	return err
}
