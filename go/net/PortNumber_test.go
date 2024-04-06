package net

import (
	"testing"
)

func TestPortNumber_FromInt(t *testing.T) {
	var n PortNumber
	for i := int(0); i < 65535; i++ {
		if err := n.FromInt(i); err != nil {
			t.Fatalf("i:%d ->%v", i, err)
		}
	}
	if err := n.FromInt(int(-1)); err == nil && err.Error() != "invalid number must be 0...65535" {
		t.Fatal("expected overflow error")
	}
	if err := n.FromInt(int(65536)); err == nil && err.Error() != "invalid number must be 0...65535" {
		t.Fatal("expected overflow error")
	}
}
