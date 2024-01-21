package vmNet

import "testing"

func TestNetworkInterfaceDescriptor_Type(t *testing.T) {
	var i NetworkInterfaceDescriptor
	if i.interfaceType != Undefined {
		t.Fatal("initial state should be undefined")
	}
	if err := i.Type(Bridge); err != nil {
		t.Fatal("unexpected error")
	}
	if i.interfaceType != Bridge {
		t.Fatal("Failed to set interface type (bridge)")
	}
	if err := i.Type(NAT); err != nil {
		t.Fatal("unexpected error")
	}
	if i.interfaceType != NAT {
		t.Fatal("Failed to set interface type (NAT)")
	}
}
