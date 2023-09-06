package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// boundsCheck - Given an integer value, return error if outside of min/max bounds
func boundsCheck(value int) (err error) {
	if (value < minCacheLevel) || (value > maxCacheLevel) {
		err = fmt.Errorf(errors.IndexOutOfRange)
	}
	return err
}
