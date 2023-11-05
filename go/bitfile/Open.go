package bitfile

/*
 * CRSCE bitfile
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer open method
 */
import (
	"fmt"
	"os"
)

// Open - Open the named file
func (o *BitFile) Open(fileName *string) (err error) {
	o.file, err = os.Open(*fileName)
	if err != nil {
		return fmt.Errorf("cannot open file (%s): %v", *fileName, err)
	} else {
		o.buffer = make([]byte, bufferSize)
		o.bufferSize, err = o.file.Read(o.buffer)
	}
	return err
}
