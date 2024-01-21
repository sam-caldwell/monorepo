package virtualization

import "testing"

func TestVM_GetImage(t *testing.T) {
	var vm VM
	if i := vm.GetImage(); i != "" {
		t.Fatalf("initial image name mismatch")
	}
	vm.image = "test"
	if i := vm.GetImage(); i != "test" {
		t.Fatalf("Image name mismatch")
	}
}
