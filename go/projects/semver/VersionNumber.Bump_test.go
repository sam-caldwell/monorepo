package semver

import "testing"

func TestVersionNumber_Bump(t *testing.T) {
	var v VersionNumber

	//Get the maximum value of Version number using bitwise math...much faster.
	maxValue := VersionNumber(2<<(versionNumberSize-1) - 1)

	if v != 0 {
		t.Fail()
	}

	if err := v.Bump(); err != nil {
		t.Fatal(err)
	}

	if v != 1 {
		t.Fatal("Increment failed")
	}

	v = maxValue
	if err := v.Bump(); err == nil {
		t.Fatal(err)
	}
}
