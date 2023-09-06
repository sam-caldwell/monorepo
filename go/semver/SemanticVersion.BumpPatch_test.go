package semver

import "testing"

func TestSemanticVersion_BumpPatch(t *testing.T) {
	var v SemanticVersion

	//Get the maximum value of Version number using bitwise math...much faster.
	maxValue := VersionNumber(2<<(versionNumberSize-1) - 1)

	if v.minor != 0 {
		t.Fatal("initial state: minor")
	}
	if v.patch != 0 {
		t.Fatal("initial state: patch")
	}
	if err := v.BumpPatch(); err != nil {
		t.Fatal(err)
	}
	if v.patch != 1 {
		t.Fatal("increment 1")
	}
	v.patch = maxValue
	if err := v.BumpPatch(); err != nil {
		if v.minor != 1 {
			t.Fatal("increment 1 (minor)")
		}
	}
}
