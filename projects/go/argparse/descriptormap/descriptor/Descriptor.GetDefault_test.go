package descriptor

import (
	"testing"
)

func TestDescriptor_GetDefault(t *testing.T) {
	var arg Descriptor
	arg.dValue = "foo"

	if typ := arg.dValue; typ != "foo" {
		t.Fatalf("Descriptor.GetDefault() failed.  Got: %s", typ)
	}
}
