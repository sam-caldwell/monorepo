package projectmanifest

/*
 * projects/repotool/manifest/Manifest.IsBuildEnabled.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines IsBuildEnabled() will return the
 * internal state of Options.BuildEnabled
 */

func (manifest *Manifest) IsBuildEnabled() bool {
	return manifest.Options.BuildEnabled
}
