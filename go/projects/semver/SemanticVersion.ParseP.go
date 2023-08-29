package semver

import (
	"strings"
)

// ParseP - Parse a string into the semantic version numeric form
func (ver *SemanticVersion) ParseP(v *string) (err error) {

	parts := strings.Split(*v, dotDelimiter)

	if len(parts) >= 1 {
		parts[0] = strings.TrimLeft(parts[0], "v")
		if err := ver.major.PutP(&parts[0]); err != nil {
			return err
		}
	}
	if len(parts) >= 2 {
		if err := ver.minor.PutP(&parts[1]); err != nil {
			return err
		}
	}
	if len(parts) >= 3 {
		if err := ver.patch.PutP(&parts[2]); err != nil {
			return err
		}
	}
	return nil
}
