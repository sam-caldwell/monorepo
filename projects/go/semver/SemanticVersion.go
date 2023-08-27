package semver

// SemanticVersion - A semantic version (v0.0.0) with methods
type SemanticVersion struct {
	major VersionNumber

	minor VersionNumber

	patch VersionNumber
}
