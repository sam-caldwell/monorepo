package semver

import (
	"fmt"
	"testing"
)

func TestVersionNumber_PutP(t *testing.T) {
	var v VersionNumber
	var s string

	//Get the maximum value of Version number using bitwise math...much faster.
	maxValue := VersionNumber(2<<(versionNumberSize-1) - 1)

	s = "1"
	if err := v.PutP(&s); err != nil {
		t.Fatal("put failed on 1")
	}

	s = fmt.Sprintf("%d", maxValue)
	if err := v.PutP(&s); err != nil {
		t.Fatal("put failed on max")
	}

	s = "a"
	if err := v.PutP(&s); err == nil {
		t.Fatal("put failed on 1")
	}
}
