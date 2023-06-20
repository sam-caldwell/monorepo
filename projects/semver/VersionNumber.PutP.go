package semver

import "strconv"

// PutP - Convert string to uint (VersionNumber) and store
func (ver *VersionNumber) PutP(s *string) error {
	n, err := strconv.ParseUint(*s, versionNumberBase, versionNumberSize)
	if err != nil {
		return err
	}
	*ver = VersionNumber(n)
	return nil
}
