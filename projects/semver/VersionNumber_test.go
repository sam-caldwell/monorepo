package semver

import (
	"testing"
)

func TestVersionNumber_Type(t *testing.T) {

	const minValue = 0

	//Get the maximum value of Version number using bitwise math...much faster.
	maxValue := VersionNumber(2<<(versionNumberSize-1) - 1)
	t.Log(maxValue)

	var v VersionNumber
	v = maxValue - 1
	if v++; v != maxValue {
		t.Fail()
	}
	if v++; v != minValue {
		t.Fatal("Expected 8bit unsigned (overflow)")
	}
	v = 1
	if v--; v != minValue {
		t.Fail()
	}
	if v--; v != maxValue {
		t.Fatal("Expected 8bit unsigned (underflow)")
	}
}
