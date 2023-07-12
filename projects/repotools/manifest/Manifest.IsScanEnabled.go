package projectmanifest

/*
 * projects/repotool/manifest/IsSignEnabled.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines IsSignEnabled() will return the
 * internal state of Options.SignEnabled
 */

func (manifest *Manifest) IsScanEnabled() bool {
	return manifest.Options.ScanEnabled
}
