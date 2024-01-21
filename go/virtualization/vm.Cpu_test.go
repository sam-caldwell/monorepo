package virtualization

import "testing"

func TestVmCpu(t *testing.T) {
	var vm VM
	if vm.cpuCount != 0 {
		t.Fatalf("initial count expected 0")
	}
	n := vm.Cpu(1)
	if n != &vm {
		t.Fatalf("expect VM.Cpu() to return the address of the VM instance")
	}
	if vm.cpuCount != 1 {
		t.Fatalf("expect cpuCount to be 1 after setting it with Cpu() method")
	}

}
