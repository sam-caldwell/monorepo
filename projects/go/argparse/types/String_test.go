package types

import (
	"testing"
)

func TestArgTypes_String(t *testing.T) {
	expected := []string{"Boolean", "Flag", "Float", "Integer", "String"}
	for pos, thisType := range []ArgTypes{Boolean, Flag, Float, Integer, String} {
		if thisType.String() != expected[pos] {
			t.Fatalf("type mismatch: '%s' != '%s'", thisType.String(), expected)
		}
	}
	var thisOne ArgTypes = 99
	if thisOne.String() != "unknown-type(99)" {
		t.Fatal("unknown type .String() result mismatch")
	}
}
