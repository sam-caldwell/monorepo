package projectmanifest

/*
 * projects/repotool/manifest/Error.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines Error() which returns the
 * internal error state and is intended to terminate
 * any chain of Manifest methods.
 */

// Error - return the internal error state
func (manifest *Manifest) Error() error {
	return manifest.err
}
