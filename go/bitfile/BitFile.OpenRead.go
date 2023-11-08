package bitfile

/*
 * CRSCE bitfile
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer open method
 */
import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
)

// OpenRead - OpenRead the named file (error if not exists)
func (o *BitFile) OpenRead(fileName *string) (err error) {
	o.file, err = os.Open(*fileName)
	if err != nil {
		return fmt.Errorf("cannot open file (%s): %v", *fileName, err)
	} else {
		o.buffer = make([]byte, bufferSize)
		o.bufferSize, err = o.file.Read(o.buffer)
	}
	return err
}

// Create - Create or open a bit file
func (o *BitFile) Create(fileName *string) (err error) {
	if !file.Existsp(fileName) {
		o.file, err = os.Create(*fileName)
	} else {
		err = o.OpenRead(fileName)
	}
	return err
}
