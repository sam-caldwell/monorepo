package descriptor

import "testing"

func TestDescriptor_GetPosition(t *testing.T) {
	var arg Descriptor
	if arg.pos != 0 {
		t.Fatal("initial state failed")
	}
	if arg.GetPosition() != 0 {
		t.Fatal("initial state returned wrong")
	}
	arg.pos = 99999
	if arg.GetPosition() != 99999 {
		t.Fatal("Descriptor.GetPosition() returns wrong values.")
	}
}
