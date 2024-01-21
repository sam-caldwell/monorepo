package virtualization

import (
	"github.com/sam-caldwell/monorepo/go/virtualization/vmNet"
	"testing"
)

func TestVM_NetworkInterface(t *testing.T) {
	t.Run("readonly protections", func(t *testing.T) {
		var vm VM
		if vm.readOnly {
			t.Fatalf("expect readonly false")
		}
		vm.readOnly = true
		if !vm.readOnly {
			t.Fatalf("expect readonly true")
		}
		if p := vm.NetworkInterface("eth0", vmNet.Bridge); p != &vm {
			t.Fatalf("expect address of vm object")
		}
		if vm.errors.Size() == 0 {
			t.Fatalf("expect an error on readonly violation")
		}
		if err := vm.errors.Dump(); err[0].Error() != "readonly violation (NetworkInterface)" {
			t.Fatalf("error message mismatch")
		}
	})
	t.Run("initial state", func(t *testing.T) {
		var vm VM
		if len(vm.networkInterfaces) != 0 {
			t.Fatalf("expected no network interfaces")
		}
		if vm.errors.Size() != 0 {
			t.Fatalf("expect no errors")
		}
	})
	t.Run("happy path: create one interface", func(t *testing.T) {
		var vm VM
		vm.NetworkInterface("eth0", vmNet.Bridge).
			NetworkInterface("eth1", vmNet.NAT).
			NetworkInterface("eth2", vmNet.Bridge)

		if vm.errors.Size() != 0 {
			t.Fatalf("expected zero errors")
		}
		if len(vm.networkInterfaces) != 3 {
			t.Fatalf("expected 3 interfaces")
		}
		if i := vm.GetNetworkInterface(0); i.GetName() != "eth0" {
			t.Fatalf("interface name mismatch 0")
		}
		if i := vm.GetNetworkInterface(1); i.GetName() != "eth1" {
			t.Fatalf("interface name mismatch 1")
		}
		if i := vm.GetNetworkInterface(2); i.GetName() != "eth2" {
			t.Fatalf("interface name mismatch 2")
		}
		if i := vm.GetNetworkInterface(0); i.GetType() != vmNet.Bridge {
			t.Fatalf("interface type mismatch 0")
		}
		if i := vm.GetNetworkInterface(1); i.GetType() != vmNet.NAT {
			t.Fatalf("interface type mismatch 1")
		}
		if i := vm.GetNetworkInterface(2); i.GetType() != vmNet.Bridge {
			t.Fatalf("interface type mismatch 2")
		}
	})
	t.Run("sad path: bad interface name", func(t *testing.T) {
		var vm VM
		_ = vm.NetworkInterface("bad.eth0", vmNet.Bridge)
		if vm.errors.Size() < 1 {
			t.Fatalf("expected one error")
		}
		_ = vm.NetworkInterface("bad.eth1", vmNet.Bridge)
		if vm.errors.Size() < 2 {
			t.Fatalf("expected two errors")
		}
		if i := vm.errors.Dump(); i[0].Error() != "invalid name" {
			t.Fatalf("error mismatch: %s", i[0].Error())
		}
	})
}
