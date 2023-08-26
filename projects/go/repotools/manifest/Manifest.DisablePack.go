package projectmanifest

/*
 * projects/repotool/manifest/Manifest.DisablePack.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the DisablePack() method
 * which will set the PackEnabled option to false.
 */

// DisablePack - Set Options.PackEnabled to false
func (manifest *Manifest) DisablePack() *Manifest {
	manifest.Options.PackEnabled = false
	return manifest
}
