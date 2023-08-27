package projectmanifest

/*
 * projects/repotool/manifest/Manifest.EnableScan.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the EnableScan() method
 * which will set the ScanEnabled option to true.
 */

// EnableScan - Set Options.ScanEnabled to true
func (manifest *Manifest) EnableScan() *Manifest {
	manifest.Options.ScanEnabled = true
	return manifest
}
