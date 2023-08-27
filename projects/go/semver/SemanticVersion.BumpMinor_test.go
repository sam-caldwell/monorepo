package semver

import "testing"

func TestSemanticVersion_BumpMinor(t *testing.T) {
	var v SemanticVersion

	//Get the maximum value of Version number using bitwise math...much faster.
	maxValue := VersionNumber(2<<(versionNumberSize-1) - 1)

	if v.major != 0 {
		t.Fail()
	}
	if v.minor != 0 {
		t.Fail()
	}
	if err := v.BumpMinor(); err != nil {
		t.Fatal(err)
	}
	if v.minor != 1 {
		t.Fail()
	}
	v.minor = maxValue
	if err := v.BumpMinor(); err != nil {
		if v.major != 1 {
			t.Fail()
		}
	}
}
