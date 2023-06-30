package systemrecon

/*
 * CpuCache
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file will call the system-specific getCacheSizes() function
 * to query the system for the L1, L2 and L3 CPU Cache size (in KB)
 * and this function will return the same as a colon-delimited string.
 */
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
