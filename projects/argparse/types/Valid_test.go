package types

import (
	"testing"
)

func TestArgTypes_Valid(t *testing.T) {
	for _, thisType := range []ArgTypes{Boolean, Flag, Float, Integer, String} {
		if err := thisType.Valid(); err != nil {
			t.Fatalf("type is invalid (should be):%s", thisType.String())
		}
	}
	var thisType ArgTypes = 99
	if err := thisType.Valid(); err == nil {
		t.Fatalf("type is valid (should not be):%s", thisType.String())
	}
}
