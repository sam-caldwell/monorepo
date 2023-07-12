package projectmanifest

/*
 * projects/repotool/manifest/IsSignEnabled.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines IsSignEnabled() will return the
 * internal state of Options.SignEnabled
 */

// EnableSign - Set Options.SignEnabled to true
func (manifest *Manifest) EnableSign() *Manifest {
	manifest.Options.SignEnabled = true
	return manifest
}
