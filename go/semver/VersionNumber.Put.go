package semver

// Put - Convert string to uint (VersionNumber) and store
func (ver *VersionNumber) Put(s string) error {
	return ver.PutP(&s)
}
