package bitfile

/*
 * Reader.Open() method
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Open an existing file or throw an error if not found.
 */
import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"os"
)

// Open - Open the named file (error if not exists)
func (o *Reader) Open(fileName *string, blockSize uint) (err error) {

	if blockSize < MinimumBlockSize {
		return fmt.Errorf(errors.ValueTooSmall)
	}
	o.blockSize = blockSize
	if o.file, err = os.Open(*fileName); err != nil {
		err = fmt.Errorf(errors.CannotOpenFile+errors.Details+errors.Details, *fileName, err)
	}
	return err

}
