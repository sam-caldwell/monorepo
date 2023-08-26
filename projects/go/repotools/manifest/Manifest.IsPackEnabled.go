package projectmanifest

/*
 * projects/repotool/manifest/Manifest.IsPackEnabled.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines IsPackEnabled() will return the
 * internal state of Options.SignEnabled
 */

func (manifest *Manifest) IsPackEnabled() bool {
	return manifest.Options.PackEnabled
}
