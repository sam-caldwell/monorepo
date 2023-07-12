package projectmanifest

/*
 * projects/repotool/manifest/Manifest.SupportsArch.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file searches the Manifest SupportedPlatforms
 * and returns true if the Manifest supports the given
 * CPU Architecture.
 */

// SupportsArch  - Evaluate if the Manifest supports the given CPU architecture
func (manifest *Manifest) SupportsArch(arch string) bool {
	for _, platform := range manifest.SupportedPlatforms {
		if platform.CpuArch == arch {
			return true
		}
	}
	return false
}
