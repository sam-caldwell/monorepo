package vmNet

import "testing"

func TestNetworkInterfaceDescriptor_Name(t *testing.T) {
	var net NetworkInterfaceDescriptor
	if net.name != "" {
		t.Fatal("expect net.name to be empty initially")
	}
	if err := net.Name("eth0"); err != nil {
		t.Fatal("expect no error")
	}
	if err := net.Name("bad.interface"); err == nil {
		t.Fatal("expected error")
	}
}
