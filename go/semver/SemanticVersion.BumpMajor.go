package semver

// BumpMajor - bump version part
func (ver *SemanticVersion) BumpMajor() error {
	return ver.major.Bump()
}
