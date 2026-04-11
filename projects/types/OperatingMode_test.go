package types

import "testing"

func TestOperatingMode_FromString(t *testing.T) {
	var o OperatingMode
	if o != 0 {
		t.Fatal("Undefined")
	}
	if err := o.FromString("client"); (err != nil) && (o != 1) {
		t.Fatal("client should be 1")
	}
	if err := o.FromString("server"); (err != nil) && (o != 2) {
		t.Fatal("server should be 2")
	}
	if err := o.FromString("bad"); err == nil {
		t.Fatal("bad input should throw error")
	}
}

func TestOperatingMode_ToString(t *testing.T) {
	var o OperatingMode
	if o != 0 {
		t.Fatal("Undefined")
	}
	o = 1
	if s := o.ToString(); s != "client" {
		t.Fatal("expected client")
	}
	o = 2
	if s := o.ToString(); s != "server" {
		t.Fatal("expected server")
	}
}
