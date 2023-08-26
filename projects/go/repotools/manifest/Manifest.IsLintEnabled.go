package projectmanifest

/*
 * projects/repotool/manifest/Manifest.IsLintEnabled.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines IsLintEnabled() will return the
 * internal state of Options.LintEnabled
 */

// IsLintEnabled - return the state of Options.LintEnabled
func (manifest *Manifest) IsLintEnabled() bool {
	return manifest.Options.LintEnabled
}
