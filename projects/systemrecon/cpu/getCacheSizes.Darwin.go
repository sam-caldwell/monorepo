//go:build darwin
// +build darwin

package systemrecon

import (
	runcommand "github.com/sam-caldwell/go/v2/projects/RunCommand"
	"strconv"
	"strings"
)

func getCacheSizes(executor runcommand.CommandExecutor) (l1, l2, l3 int, err error) {
	cacheLevels := []string{"hw.l1icachesize", "hw.l2cachesize", "hw.l3cachesize"}
	results := []int{0, 0, 0}

	for i, level := range cacheLevels {
		out, err := executor.Execute("sysctl", "-n", level)
		if err != nil {
			return 0, 0, 0, err
		}

		size, err := strconv.Atoi(strings.TrimSpace(string(out)))
		if err != nil {
			return 0, 0, 0, err
		}

		results[i] = size / 1024
	}

	return results[0], results[1], results[2], nil
}
