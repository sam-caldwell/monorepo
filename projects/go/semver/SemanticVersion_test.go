package semver

import "testing"

func TestSemanticVersion_struct(t *testing.T) {
	var v SemanticVersion
	if v.major != 0 {
		t.Fail()
	}
	if v.minor != 0 {
		t.Fail()
	}
	if v.patch != 0 {
		t.Fail()
	}
}
