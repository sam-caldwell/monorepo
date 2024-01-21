package vmNet

import "testing"

func TestNetworkInterfaceDescriptor_GetName(t *testing.T) {
	var net NetworkInterfaceDescriptor
	net.name = "test"
	if n := net.GetName(); n != "test" {
		t.Fatal("GetName() mismatch")
	}
}
