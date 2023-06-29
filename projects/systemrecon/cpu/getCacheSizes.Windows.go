//go:build windows
// +build windows

package systemrecon

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// getCacheSizes - Return a given CPU cache (L1, L2, L3)
func getCacheSizes(level int) (size int, err error) {
	if level < 1 || level > 3 {
		return 0, fmt.Errorf("invalid cache level. Must be 1, 2, or 3")
	}

	cmd := exec.Command("wmic", "cpu", "get", fmt.Sprintf("L%dCacheSize", level), "/value")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	result := strings.Split(string(output), "=")
	if len(result) != 2 {
		return 0, fmt.Errorf("failed to retrieve cache size")
	}

	cacheSizeStr := strings.TrimSpace(result[1])
	cacheSize, err := strconv.Atoi(cacheSizeStr)
	if err != nil {
		return 0, err
	}

	return cacheSize, nil
}
