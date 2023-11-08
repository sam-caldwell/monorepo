package file

/*
 * BitFile.Create() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create a file (if not found) or open file (if exists).
 */

import (
	"os"
)

// Create - Create a file if not found, or open an existing file
func (o *BitFile) Create(fileName *string) (err error) {
	if Existsp(fileName) {
		err = o.Open(fileName)
	} else {
		o.file, err = os.Create(*fileName)
	}
	return err
}
