package projectmanifest

/*
 * projects/repotool/manifest/Manifest.EnableBuild.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the EnableBuild() method
 * which will set the BuildEnabled option to true.
 */

// EnableBuild - Set Options.BuildEnabled to true
func (manifest *Manifest) EnableBuild() *Manifest {
	manifest.Options.BuildEnabled = true
	return manifest
}
