package bitfile

/*
 * Writer.Close() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Close the file for the Writer object.
 */

// Close - Close method to close the file behind the writer object.
func (o *Writer) Close() (err error) {

	if o.file != nil {
		err = o.file.Close()
	}

	return err

}
