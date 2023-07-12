package projectmanifest

/*
 * projects/repotool/manifest/DisableScan.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the DisableScan() method
 * which will set the ScanEnabled option to false.
 */

// DisableScan - Set Options.ScanEnabled to false
func (manifest *Manifest) DisableScan() *Manifest {
	manifest.Options.ScanEnabled = false
	return manifest
}
