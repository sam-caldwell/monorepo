package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/convert"
)

// CpuCache - Return the CPU cache (in KB)
func CpuCache() (result string, err error) {
	var l1, l2, l3 string
	if l1, err = convert.IntErrToStringErr(getCacheSizes(1)); err == nil {
		if l2, err = convert.IntErrToStringErr(getCacheSizes(2)); err == nil {
			l3, err = convert.IntErrToStringErr(getCacheSizes(3))
		}
	}
	return fmt.Sprintf("%s:%s:%s", l1, l2, l3), err
}
