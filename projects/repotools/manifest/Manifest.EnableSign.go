package projectmanifest

/*
 * projects/repotool/manifest/EnableSign.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the EnableSign() method
 * which will set the SignEnabled option to true.
 */

// EnableSign - Set Options.SignEnabled to false
func (manifest *Manifest) EnableSign() *Manifest {
	manifest.Options.SignEnabled = true
	return manifest
}
