package projectmanifest

/*
 * projects/repotool/manifest/Manifest.EnableSign.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines EnableSign() will set the
 * internal state of Options.SignEnabled to true
 */

// EnableSign - Set Options.SignEnabled to true
func (manifest *Manifest) EnableSign() *Manifest {
	manifest.Options.SignEnabled = true
	return manifest
}
