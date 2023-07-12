package projectmanifest

// SupportsArch  - Return boolean answer to the question "Does manifest support the given CPU Architecture?"
func (manifest *Manifest) SupportsArch(arch string) bool {
	for _, platform := range manifest.SupportedPlatforms {
		if platform.CpuArch == arch {
			return true
		}
	}
	return false
}
