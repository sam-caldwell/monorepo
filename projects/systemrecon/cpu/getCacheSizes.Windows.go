//go:build windows
// +build windows

package systemrecon

import (
	"errors"
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/convert"
	"github.com/sam-caldwell/go/v2/projects/wmi"
)

const (
	invalidCacheSz = -1
)

type Win32_CacheMemory struct {
	Purpose        string
	BlockSize      uint64
	NumberOfBlocks uint64
}

// getCacheSizes - Return a given CPU cache (L1, L2, L3)
func getCacheSizes(level int) (size int, err error) {
	cacheLevels := map[int]string{1: "L1 Cache", 2: "L2 Cache", 3: "L3 Cache"}
	cachePurpose, ok := cacheLevels[level]
	if !ok {
		return invalidCacheSz, fmt.Errorf("Invalid cache level: %d", level)
	}

	var dst []Win32_CacheMemory
	query := fmt.Sprintf("SELECT * FROM Win32_CacheMemory WHERE Purpose = '%s'", cachePurpose)
	err = wmi.Query(query, &dst)
	if err != nil {
		return invalidCacheSz, err
	}

	if len(dst) == 0 {
		return invalidCacheSz, errors.New("No data returned from WMI Query")
	}

	cacheSizeBytes := dst[0].BlockSize * dst[0].NumberOfBlocks
	return convert.BytesToKilobytes(int(cacheSizeBytes)), nil
}
