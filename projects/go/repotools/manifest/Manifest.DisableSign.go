package projectmanifest

/*
 * projects/repotool/manifest/Manifest.DisableSign.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the DisableSign() method
 * which will set the SignEnabled option to false.
 */

// DisableSign - Set Options.SignEnabled to false
func (manifest *Manifest) DisableSign() *Manifest {
	manifest.Options.SignEnabled = false
	return manifest
}
