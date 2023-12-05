package bitfile

/*
 * Reader.Close() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Close the current reader file handle.
 */

// Close - Close the current Reader
func (o *Reader) Close() (err error) {

	if o.file != nil {
		err = o.file.Close()
		if err == nil {
			o.file = nil
		}
	}
	return err
}
