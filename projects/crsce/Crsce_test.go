package crsce

import "testing"

func TestCrsce(t *testing.T) {
	var crsce Crsce
	if len(crsce.lsm) != 0 {
		t.Fatal("LSM should be empty unless properly initialized by New<n>Compressor()")
	}
	if len(crsce.vsm) != 0 {
		t.Fatal("VSM should be empty unless properly initialized by New<n>Compressor()")
	}
}
