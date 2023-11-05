package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * ByteCounter Adder Function
 */

func (c *ByteCounter) Add(value int) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	for i := 0; i < value; i++ {
		c.FastIncrement()
	}
	return err
}
