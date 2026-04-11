package hijack

import (
	"reflect"
	"testing"
)

func TestAppliedPatch(t *testing.T) {
	// Prepare test data
	originalBytes := []byte{0x90, 0x90, 0x90} // Example original bytes
	targetFunc := reflect.ValueOf(func() {})
	imposterFunc := reflect.ValueOf(func() {})

	// Create an AppliedPatch instance
	patch := AppliedPatch{
		originalBytes: originalBytes,
		target:        &targetFunc,
		imposter:      &imposterFunc,
	}

	// Assert the fields of the AppliedPatch instance
	if !reflect.DeepEqual(patch.originalBytes, originalBytes) {
		t.Fatalf("AppliedPatch originalBytes field does not match")
	}
	if !reflect.DeepEqual(patch.target, &targetFunc) {
		t.Fatalf("AppliedPatch target field does not match")
	}
	if !reflect.DeepEqual(patch.imposter, &imposterFunc) {
		t.Fatalf("AppliedPatch imposter field does not match")
	}
}
