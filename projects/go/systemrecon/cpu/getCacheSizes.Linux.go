//go:build linux
// +build linux

package systemrecon

/*
 * getCacheSizes() - Linux
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Gather the cache size information for a given cache level (1-3)
 * in KB by reading from /sys/devices/system/cpu/cpu0/cache/index{0,1,2,3}/size
 *
 * See CpuCache.md
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/wrappers/os"
	"strconv"
	"strings"
)

// getCacheSizes - Return a given CPU cache (0:0:<cache_size>)
func getCacheSizes(level int) (size int, err error) {
	const (
		cacheFile = "/sys/devices/system/cpu/cpu0/cache/index%d/size"
	)
	var raw []byte

	if err = boundsCheck(level); err == nil {
		if raw, err = os.ReadFile(fmt.Sprintf(cacheFile, level)); err == nil {
			if size, err = strconv.Atoi(strings.TrimSuffix(string(raw), "K")); err == nil {
				return size, err
			}
		}
	}
	return invalidCacheSz, err
}
