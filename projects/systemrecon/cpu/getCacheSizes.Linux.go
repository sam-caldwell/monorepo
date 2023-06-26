//go:build linux
// +build linux

package systemrecon

import (
	"fmt"
	runcommand "github.com/sam-caldwell/go/v2/projects/RunCommand"
	"github.com/sam-caldwell/go/v2/projects/convert"
	"strconv"
	"strings"
)

// getCacheSizes - Return a given CPU cache (L1, L2, L3)
func getCacheSizes(executor runcommand.CommandExecutor, level int) (size int, err error) {
	cacheLevels := []string{"cache size", "cache size", "cache size"} // Linux does not distinguish between L1, L2, L3 cache sizes
	if level >= 0 && level <= len(cacheLevels) {
		var raw []byte
		if raw, err = executor.Execute("grep", cacheLevels[level-1], "/proc/cpuinfo"); err == nil {
			lines := strings.Split(strings.TrimSpace(string(raw)), "\n")
			lastLine := lines[len(lines)-1] // The last line contains the L3 cache size
			parts := strings.Fields(lastLine)
			if len(parts) > 0 {
				if size, err = strconv.Atoi(parts[0]); err == nil {
					return convert.BytesToKilobytes(size), nil
				}
			}
		}
	} else {
		err = fmt.Errorf("Invalid cache level: %d", level)
	}
	return invalidCacheSz, err
}
