package semver

import "fmt"

func (ver *VersionNumber) Bump() error {

	//Get the maximum value of Version number using bitwise math...much faster.
	maxValue := VersionNumber(2<<(versionNumberSize-1) - 1)

	if *ver == maxValue {
		return fmt.Errorf(errVersionOverflow)
	}
	*ver++
	return nil
}
