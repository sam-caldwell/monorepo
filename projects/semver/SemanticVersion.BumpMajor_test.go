package semver

import "testing"

func TestSemanticVersion_BumpMajor(t *testing.T) {
	var v SemanticVersion

	//Get the maximum value of Version number using bitwise math...much faster.
	maxValue := VersionNumber(2<<(versionNumberSize-1) - 1)

	if v.major != 0 {
		t.Fail()
	}
	if err := v.BumpMajor(); err != nil {
		t.Fatal(err)
	}
	if v.major != 1 {
		t.Fail()
	}
	v.major = maxValue
	if err := v.BumpMajor(); err == nil {
		t.Fatal(err)
	}
}
