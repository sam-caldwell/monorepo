package projectmanifest

/*
 * projects/repotool/manifest/DisableBuild.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the DisableBuild() method
 * which will set the BuildEnabled option to false.
 */

// DisableBuild - Set Options.BuildEnabled to false
func (manifest *Manifest) DisableBuild() *Manifest {
	manifest.Options.BuildEnabled = false
	return manifest
}
