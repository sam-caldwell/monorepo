package projectmanifest

/*
 * projects/repotool/manifest/Manifest.SetName.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines SetName() set the project name
 */

// SetName - set project name
func (manifest *Manifest) SetName(name string) *Manifest {
	if manifest.err != nil {
		return manifest
	}
	manifest.Name = name
	return manifest
}
