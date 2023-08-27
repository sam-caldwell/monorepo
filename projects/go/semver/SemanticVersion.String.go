package semver

import "fmt"

// String - return a semantic version string.
func (ver *SemanticVersion) String() string {
	return fmt.Sprintf(tagFormat, ver.major, ver.minor, ver.patch)
}
