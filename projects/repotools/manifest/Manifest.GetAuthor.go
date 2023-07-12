package projectmanifest

/*
 * projects/repotool/manifest/Manifest.GetAuthor.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines GetAuthor which will return
 * the name of the project's author.
 */

// GetAuthor - Return the project author's name
func (manifest *Manifest) GetAuthor() string {
	return manifest.Author
}
