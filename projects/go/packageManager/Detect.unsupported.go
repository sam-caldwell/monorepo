//go:build !windows && !linux && !darwin
// +build !windows,!linux,!darwin

package packageManager

func Detect() (pkg pkgManagerFunc, err error) {
	return nil, nil
}
