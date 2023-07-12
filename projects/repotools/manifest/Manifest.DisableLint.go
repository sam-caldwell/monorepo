package projectmanifest

/*
 * projects/repotool/manifest/DisableLint()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the DisableLint() method
 * which will set the DisableLint option to false.
 */

// DisableLint - Set Options.DisableLint to false
func (manifest *Manifest) DisableLint() *Manifest {
	manifest.Options.LintEnabled = false
	return manifest
}
