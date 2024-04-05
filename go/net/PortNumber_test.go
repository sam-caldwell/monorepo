package net

import (
	"testing"
)

func TestPortNumber_FromInt(t *testing.T) {
	var n PortNumber
	for i := uint(0); i < 65535; i++ {
		if err := n.FromInt(i); err != nil {
			t.Fatalf("i:%d ->%v", i, err)
		}
	}
	if err := n.FromInt(uint(65536)); err == nil && err.Error() != "invalid number must be 0...65535" {
		t.Fatal("expected overflow error")
	}
}
