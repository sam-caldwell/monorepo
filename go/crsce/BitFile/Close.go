package bitfile

/*
 * CRSCE BitFile
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer close method
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
)

func (o *BitFile) Close() {
	if o.file != nil {
		err := o.file.Close()
		exit.OnError(err, exit.GeneralError, fmt.Sprintf("error closing file (%s)", o.file.Name()))
		o.file = nil
	}
}
