//go:build darwin
// +build darwin

package systemrecon

import (
	"fmt"
	runcommand "github.com/sam-caldwell/go/v2/projects/RunCommand"
	"github.com/sam-caldwell/go/v2/projects/convert"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"strconv"
	"strings"
)

const (
	invalidCacheSz = -1
)

// getCacheSizes - Return a given CPU cache (L1, L2, L3)
func getCacheSizes(executor runcommand.CommandExecutor, level int) (size int, err error) {
	cacheLevels := []string{"hw.l1icachesize", "hw.l2cachesize", "hw.l3cachesize"}
	if level >= 0 && level <= len(cacheLevels) {
		var raw []byte
		if raw, err = executor.Execute("sysctl", "-n", cacheLevels[level-1]); err == nil {
			if size, err := strconv.Atoi(strings.TrimSpace(string(raw))); err == nil {
				return convert.BytesToKilobytes(size), err
			}
		}
	} else {
		err = fmt.Errorf(errors.IndexOutOfRange)
	}
	return invalidCacheSz, err
}
