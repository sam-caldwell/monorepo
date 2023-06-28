//go:build windows
// +build windows

package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/convert"
	"strconv"
	"strings"
)

// getCacheSizes - Return a given CPU cache (L1, L2, L3)
func getCacheSizes(level int) (size int, err error) {
	cacheLevels := map[int]string{1: "L1 Cache", 2: "L2 Cache", 3: "L3 Cache"}
	cachePurpose, ok := cacheLevels[level]
	if !ok {
		return invalidCacheSz, fmt.Errorf("Invalid cache level: %d", level)
	}

	command := fmt.Sprintf("wmic CacheMemory where Purpose='%s' get BlockSize,NumberOfBlocks /value", cachePurpose)
	raw, err := executor.Execute("cmd", "/C", command)
	if err != nil {
		return invalidCacheSz, err
	}

	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")
	if len(lines) != 2 {
		return invalidCacheSz, fmt.Errorf("Unexpected output from wmic command")
	}

	blockSize, err := strconv.Atoi(strings.Split(lines[0], "=")[1])
	if err != nil {
		return invalidCacheSz, err
	}

	numberOfBlocks, err := strconv.Atoi(strings.Split(lines[1], "=")[1])
	if err != nil {
		return invalidCacheSz, err
	}

	cacheSizeBytes := blockSize * numberOfBlocks
	return convert.BytesToKilobytes(cacheSizeBytes), nil
}
