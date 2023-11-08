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

// CreateOrOpen - Create or open a bit file
func (o *BitFile) CreateOrOpen(fileName *string) (err error) {
	if !Existsp(fileName) {
		o.file, err = os.Create(*fileName)
	} else {
		err = o.OpenRead(fileName)
	}
	return err
}
