package projectmanifest

/*
 * projects/repotool/manifest/EnableLint.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the EnableLint() method
 * which will set the LintEnabled option to true.
 */

// EnableLint - Set Options.LintEnabled to false
func (manifest *Manifest) EnableLint() *Manifest {
	manifest.Options.LintEnabled = true
	return manifest
}
