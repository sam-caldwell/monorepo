package systemrecon

import (
	"fmt"
	runcommand "github.com/sam-caldwell/go/v2/projects/RunCommand"
	"github.com/sam-caldwell/go/v2/projects/convert"
)

// CpuCache - Return the CPU cache (in KB)
func CpuCache(runner runcommand.CommandExecutor) (result string, err error) {
	var l1, l2, l3 string
	if l1, err = convert.IntErrToStringErr(getCacheSizes(runner, 1)); err == nil {
		if l2, err = convert.IntErrToStringErr(getCacheSizes(runner, 2)); err == nil {
			if l3, err = convert.IntErrToStringErr(getCacheSizes(runner, 3)); err == nil {
				//Noop
			}
		}
	}
	return fmt.Sprintf("%s:%s:%s", l1, l2, l3), err
}
