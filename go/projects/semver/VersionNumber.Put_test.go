package semver

import (
	"fmt"
	"testing"
)

func TestVersionNumber_Put(t *testing.T) {
	var v VersionNumber

	//Get the maximum value of Version number using bitwise math...much faster.
	maxValue := VersionNumber(2<<(versionNumberSize-1) - 1)

	if err := v.Put("1"); err != nil {
		t.Fatal("put failed on 1")
	}

	if err := v.Put(fmt.Sprintf("%d", maxValue)); err != nil {
		t.Fatal("put failed on max")
	}

	if err := v.Put("a"); err == nil {
		t.Fatal("put failed on 1")
	}
}
