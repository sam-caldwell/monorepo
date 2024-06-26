//go:build darwin
// +build darwin

package systemrecon

import (
	"errors"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/convert"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestCpuCache(t *testing.T) {
	getExpected := func(lvl string) (size int) {
		var raw []byte
		var err error
		if raw, err = exec.Command("sysctl", "-n", lvl).Output(); err == nil {
			if size, err = strconv.Atoi(strings.TrimSpace(string(raw))); err == nil {
				return convert.BytesToKilobytes(size)
			}
		}
		t.Fatalf("setup error: %s", err)
		return
	}
	expectedCache := fmt.Sprintf("%d:%d:%d",
		getExpected("hw.l1icachesize"), getExpected("hw.l2cachesize"), getExpected("hw.l3cachesize"),
	)
	expectedErr := error(nil) // Expected error, or nil if no error is expected

	cache, err := CpuCache()

	if cache != expectedCache {
		t.Fatalf("Unexpected cache size.\n"+
			"\tExpected: %v,\n"+
			"\tGot: %v",
			expectedCache, cache)
	}

	if !errors.Is(err, expectedErr) {
		t.Fatalf("Unexpected error.\n"+
			"\tExpected: %v,\n"+
			"\tGot: %v",
			expectedErr, err)
	}
}
