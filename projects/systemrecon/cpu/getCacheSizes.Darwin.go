//go:build darwin
// +build darwin

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
	"github.com/sam-caldwell/go/v2/projects/convert"
	"os/exec"
	"strconv"
	"strings"
)

// getCacheSizes - Return a given CPU cache (L1, L2, L3)
func getCacheSizes(level int) (size int, err error) {
	var raw []byte
	cacheLevels := []string{
		"hw.l1icachesize",
		"hw.l2cachesize",
		"hw.l3cachesize",
	}
	if err = boundsCheck(level); err == nil {
		if raw, err = exec.Command("sysctl", "-n", cacheLevels[level-1]).Output(); err == nil {
			if size, err = strconv.Atoi(strings.TrimSpace(string(raw))); err == nil {
				return convert.BytesToKilobytes(size), err
			}
		}
	}
	return invalidCacheSz, err
}
