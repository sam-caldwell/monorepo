//go:build !windows && !linux && !darwin
// +build !windows,!linux,!darwin

package systemrecon

// getCacheSizes - Return a given CPU cache (L1, L2, L3)
func getCacheSizes(level int) (size int, err error) {
	return 0, nil
}
