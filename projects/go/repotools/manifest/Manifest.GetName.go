package projectmanifest

/*
 * projects/repotool/manifest/Manifest.GetName.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines GetName() which will
 * return the project's name.
 */

// GetName - return project name
func (manifest *Manifest) GetName() string {
	return manifest.Name
}
