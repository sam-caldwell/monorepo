package projectmanifest

/*
 * projects/repotool/manifest/GetAuthor.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines GetLanguage() which will
 * return the project's language.
 */

// GetLanguage - Return project language.
func (manifest *Manifest) GetLanguage() string {
	return manifest.Options.Language
}
