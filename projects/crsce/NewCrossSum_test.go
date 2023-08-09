package crsce

import "testing"

func TestNewCrossSum(t *testing.T) {
	cs := NewCrossSum()
	if cs == nil {
		t.Fatal("cs should not be nil")
	}
	if cs.sum != 0 {
		t.Fatal("cross sum expects 0")
	}
	if cs.hash == nil {
		t.Fatal("cross sum expects hash to be initialized")
	}
}
