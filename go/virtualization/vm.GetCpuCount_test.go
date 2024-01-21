package virtualization

import "testing"

func TestVmGetCpuCount(t *testing.T) {
	var vm VM
	vm.cpuCount = 1337
	if c := vm.GetCpuCount(); c != vm.cpuCount {
		t.Fatalf("GetCpuCount() does not return expected count")
	}
	if c := vm.GetCpuCount(); c != 1337 {
		t.Fatalf("GetCpuCount() does not return expected 1337")
	}
}
