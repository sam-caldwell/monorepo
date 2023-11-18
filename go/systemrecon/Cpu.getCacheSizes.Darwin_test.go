//go:build darwin
// +build darwin

package systemrecon

import (
	"errors"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/convert"
	e "github.com/sam-caldwell/monorepo/go/exit/errors"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestGetCacheSizes(t *testing.T) {
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

	// Test case 1: Valid cache level
	level := 1
	expectedSize := getExpected("hw.l1icachesize")
	expectedErr := error(nil)

	size, err := getCacheSizes(level)

	if size != expectedSize {
		t.Errorf("Unexpected cache size. Expected: %d, got: %d", expectedSize, size)
	}

	if !errors.Is(err, expectedErr) {
		t.Errorf("Unexpected error. Expected: %v, got: %v", expectedErr, err)
	}

	// Test case 2: Invalid cache level
	level = 4
	expectedSize = invalidCacheSz
	expectedErr = fmt.Errorf(e.IndexOutOfRange)

	size, err = getCacheSizes(level)

	if size != expectedSize {
		t.Errorf("Unexpected cache size. Expected: %d, got: %d", expectedSize, size)
	}

	if err.Error() != expectedErr.Error() {
		t.Errorf("Unexpected error. Expected: %v, got: %v", expectedErr, err)
	}
}
