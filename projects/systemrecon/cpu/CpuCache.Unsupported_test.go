//go:build !darwin && !windows && !linux
// +build !darwin,!windows,!linux

package systemrecon

func TestCpuCache(t *testing.T) {
	t.Fatal("test not implemented")
}
