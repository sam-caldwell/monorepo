package systemrecon

/*
 * CpuCache
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file will call the system-specific getCacheSizes() function
 * to query the system for the L1, L2 and L3 CPU Cache size (in KB)
 * and this function will return the same as a colon-delimited string.
 *
 *   L1 cache – fastest, but smallest, data and instructions
 *              This is actually two caches (data, instructions) represented as
 *              l1d and l1i respectively.
 *   L2 cache – slower, but bigger, data-only
 *   L3 cache – slowest, but biggest, data-only
 *
 * But how do we get this information?  See CpuCache.md
 */
import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/convert"
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
