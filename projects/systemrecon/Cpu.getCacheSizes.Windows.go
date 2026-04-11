//go:build windows
// +build windows

package systemrecon

/*
 * getCacheSizes() - Windows
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Gather the cache size information for a given cache level (1-3)
 * in KB by calling powershell
 *
 * See CpuCache.md
 */

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// getCacheSizes - Return a given CPU cache (L1, L2, L3)
func getCacheSizes(level int) (size int, err error) {
	const ps = `Get-WmiObject -Query \"SELECT * FROM Win32_Processor\" | Select-Object -ExpandProperty L%dCacheSize`
	var raw []byte

	if err = boundsCheck(level); err == nil {
		if raw, err = exec.Command("powershell", "/c", fmt.Sprintf(ps, level)).Output(); err == nil {
			if size, err = strconv.Atoi(strings.TrimSpace(string(raw))); err == nil {
				return size, nil
			}
		}
	}
	return invalidCacheSz, err
}
