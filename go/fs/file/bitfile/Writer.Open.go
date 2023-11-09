package bitfile

/*
 * Reader.Create() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create a file (if not found) or open file (if exists).
 */

import (
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
)

// Create - Create a file if not found, or open an existing file.
func (o *Reader) Create(fileName *string) (err error) {
	if file.Existsp(fileName) {
		err = o.Open(fileName)
	} else {
		o.file, err = os.Create(*fileName)
	}
	return err
}
