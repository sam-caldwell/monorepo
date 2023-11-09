package bitfile

/*
 * Writer.Create() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create a file to target for write operations.
 */

import (
	"os"
)

// Open - Open/create a file.  If the file exists, truncate it and open for write.
// The file mode (perms) will be set to 0666.
func (o *Writer) Open(fileName *string) (err error) {

	o.file, err = os.Create(*fileName)

	return err

}
