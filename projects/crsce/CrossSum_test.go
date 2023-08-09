package crsce

import "testing"

func TestCrossSum(t *testing.T) {
	var cs CrossSum
	if cs.sum != 0 {
		t.Fatal("cross sum expects 0")
	}
	if cs.hash != nil {
		t.Fatal("cross sum expects hash nil unless initialized by NewCrossSum()")
	}
}
