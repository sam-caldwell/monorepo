package semver

// BumpMinor - bump version part (with rollover to major)
func (ver *SemanticVersion) BumpMinor() error {
	if err := ver.minor.Bump(); err != nil {
		return ver.major.Bump()
	}
	return nil
}
