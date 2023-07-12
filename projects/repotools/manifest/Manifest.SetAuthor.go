package projectmanifest

/*
 * projects/repotool/manifest/SetAuthor.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines SetAuthor() will set the project
 * author's name in the internal state
 */

// SetAuthor - set author name
func (manifest *Manifest) SetAuthor(author string) *Manifest {
	if manifest.err != nil {
		return manifest
	}
	manifest.Author = author
	return manifest
}
