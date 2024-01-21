package virtualization

import (
	"github.com/sam-caldwell/monorepo/go/virtualization/vmNet"
	"testing"
)

func TestVM_GetNetworkInterface(t *testing.T) {
	var vm VM
	if d := vm.GetNetworkInterface(0); (d.GetName() != "") && (d.GetType() != vmNet.Undefined) {
		t.Fatalf("initial disk name state")
	}
	networkInterfaces := make([]vmNet.NetworkInterfaceDescriptor, 2)
	_ = networkInterfaces[0].Name("eth0")
	_ = networkInterfaces[0].Size(vmNet.Bridge)
	_ = networkInterfaces[1].Name("eth1")
	_ = networkInterfaces[1].Size(vmNet.NAT)
	vm.networkInterfaces = networkInterfaces

	if d := vm.GetNetworkInterface(0); (d.GetName() != "eth0") && (d.GetType() != vmNet.Bridge) {
		t.Fatalf("eth0 mismatch")
	}
	if d := vm.GetNetworkInterface(1); (d.GetName() != "eth1") && (d.GetType() != vmNet.NAT) {
		t.Fatalf("eth1 mismatch")
	}
}
