//go:build !windows && !linux && !darwin
// +build !windows,!linux,!darwin

package systemrecon

import (
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"testing"
)

func TestGetCacheSizes(t *testing.T) {
	for i := -10; i < 10; i++ {
		size, err := getCacheSizes(i)
		if size != invalidCacheSz {
			t.Fatal("Expected invalidCacheSz response")
		}
		if err == nil {
			t.Fatal("expected error")
		}
		if err.Error() != errors.UnsupportedOpsys {
			t.Fatal("expected unsupported opsys error")
		}
	}
}
