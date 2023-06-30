//go:build !windows && !linux && !darwin
// +build !windows,!linux,!darwin

package systemrecon

import (
	"testing"
)

func TestGetCacheSizes(t *testing.T) {
	t.Fatal("test not implemented")
}
