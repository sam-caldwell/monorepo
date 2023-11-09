package bitfile

/*
 * Reader.Open() method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Open an existing file or throw an error if not found.
 */
import (
	"fmt"
	"os"
)

// Open - Open the named file (error if not exists)
func (o *Reader) Open(fileName *string) (err error) {

	o.file, err = os.Open(*fileName)

	if err != nil {
		err = fmt.Errorf("cannot open file (%s): %v", *fileName, err)
	}

	return err

}
