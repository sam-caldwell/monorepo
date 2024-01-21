//go:build !darwin && !windows && !linux
// +build !darwin,!windows,!linux

package virtualization

import "testing"

func TestVm_Create(t *testing.T) {
	t.Fatal("not supported")
}
