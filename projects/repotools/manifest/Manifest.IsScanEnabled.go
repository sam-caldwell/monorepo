package projectmanifest

/*
 * projects/repotool/manifest/Manifest.IsScanEnabled.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines IsScanEnabled() will return the
 * internal state of Options.ScanEnabled
 */

func (manifest *Manifest) IsScanEnabled() bool {
	return manifest.Options.ScanEnabled
}
