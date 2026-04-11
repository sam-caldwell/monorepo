package semver

// Parse - Parse a string into the semantic version numeric form
func (ver *SemanticVersion) Parse(v string) (err error) {
	return ver.ParseP(&v)
}

func (ver *SemanticVersion) Major() VersionNumber {
	return ver.major
}

func (ver *SemanticVersion) Minor() VersionNumber {
	return ver.minor
}

func (ver *SemanticVersion) Patch() VersionNumber {
	return ver.patch
}
