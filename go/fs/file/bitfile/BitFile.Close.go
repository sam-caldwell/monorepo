package file

/*
 * BitFile.Close() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Close the current bitfile.
 */

// Close - Close the current BitFile
func (o *BitFile) Close() (err error) {
	if o.file != nil {
		err = o.file.Close()
		if err == nil {
			o.file = nil
		}
	}
	return err
}
