package projectmanifest

// SupportsOpsys - Return boolean answer to the question "Does manifest support the given operating system?"
func (manifest *Manifest) SupportsOpsys(opsys string) bool {
	for _, platform := range manifest.SupportedPlatforms {
		if platform.Opsys == opsys {
			return true
		}
	}
	return false
}
