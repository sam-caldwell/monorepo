package semver

// BumpPatch - bump version part (with rollover to minor)
func (ver *SemanticVersion) BumpPatch() error {
	if err := ver.patch.Bump(); err != nil {
		return ver.minor.Bump()
	}
	return nil
}
