package projectmanifest

/*
 * projects/repotool/manifest/Manifest.IsTestEnabled.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines IsTestEnabled() will return the
 * internal state of Options.TestEnabled
 */

func (manifest *Manifest) IsTestEnabled() bool {
	return manifest.Options.BuildEnabled
}
