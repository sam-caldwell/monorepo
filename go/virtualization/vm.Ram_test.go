package virtualization

import (
	"github.com/sam-caldwell/monorepo/go/types/size"
	"testing"
)

func TestVM_Ram(t *testing.T) {
	var vm VM
	if vm.ram != 0 {
		t.Fatal("initial state is wrong")
	}
	if p := vm.Ram(10 * size.MB); p != &vm {
		t.Fatalf("Ram() method should return the address of the current VM instance")
	}
	if vm.ram != 10*size.MB {
		t.Fatalf("state mismatch")
	}
}
