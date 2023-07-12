package projectmanifest

/*
 * projects/repotool/manifest/SupportsOpsys.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file searches the Manifest SupportedPlatforms
 * and returns true if the Manifest supports the given
 * operating system.
 */

// SupportsOpsys - Return boolean answer to the question "Does manifest support the given operating system?"
func (manifest *Manifest) SupportsOpsys(opsys string) bool {
	for _, platform := range manifest.SupportedPlatforms {
		if platform.Opsys == opsys {
			return true
		}
	}
	return false
}
