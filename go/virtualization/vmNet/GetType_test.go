package vmNet

import "testing"

func TestNetworkInterfaceDescriptor_GetType(t *testing.T) {
	var net NetworkInterfaceDescriptor
	if n := net.GetType(); n != Undefined {
		t.Fatal("expected undefined network interface type initially")
	}

	net.interfaceType = Bridge
	if n := net.GetType(); n != Bridge {
		t.Fatal("expected Bridge")
	}

	net.interfaceType = NAT
	if n := net.GetType(); n != NAT {
		t.Fatal("expected NAT")
	}

}
