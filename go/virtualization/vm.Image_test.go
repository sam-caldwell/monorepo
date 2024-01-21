package virtualization

import "testing"

func TestVM_Image(t *testing.T) {
	var vm VM
	if i := vm.image.Get(); i != "" {
		t.Fatal("unexpected initial Image state")
	}
	if p := vm.Image("imageName"); p != &vm {
		t.Fatal("VM.Image() should return pointer to the VM instance")
	}
	if i := vm.image.Get(); i != "imageName" {
		t.Fatal("image name mismatch")
	}
}
