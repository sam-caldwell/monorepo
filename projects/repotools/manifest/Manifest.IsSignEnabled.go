package projectmanifest

/*
 * projects/repotool/manifest/IsSignEnabled.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines IsSignEnabled() will return the
 * internal state of Options.SignEnabled
 */

// IsSignEnabled - return the Options.SignEnabled state
func (manifest *Manifest) IsSignEnabled() bool {
	return manifest.Options.SignEnabled
}
