package projectmanifest

/*
 * projects/repotool/manifest/EnableBuild.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the EnableBuild() method
 * which will set the BuildEnabled option to true.
 */

// EnableBuild - Set Options.BuildEnabled to false
func (manifest *Manifest) EnableBuild() *Manifest {
	manifest.Options.BuildEnabled = true
	return manifest
}
