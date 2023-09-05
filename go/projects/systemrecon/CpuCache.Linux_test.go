//go:build linux
// +build linux

package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/wrappers/os"
	"strconv"
	"strings"
	"testing"
)

func TestCpuCache(t *testing.T) {
	const cacheFile = "/sys/devices/system/cpu/cpu0/cache/index%d/size"

	for i := 1; i <= 3; i++ {
		var expected int
		if raw, err := os.ReadFile(fmt.Sprintf(cacheFile, i)); err != nil {
			t.Fatal(err)
		} else {
			if s, err := strconv.Atoi(strings.TrimSuffix(string(raw), "K")); err != nil {
				t.Fatal(err)
			} else {
				expected = s
			}
		}
		if actual, err := systemrecon.getCacheSizes(i); err != nil {
			t.Fatal(err)
		} else if actual != expected {
			t.Fatal("actual not expected")
		}
	}
	if actual, err := systemrecon.getCacheSizes(-1); err == nil {
		t.Fatal("expected error not found")
	} else if actual != -1 {
		t.Fatal("expected -1 got something else")
	}
}
