package crsce

import "testing"

func TestCrsce(t *testing.T) {
	var crsce Crsce
	if crsce.csm != nil {
		t.Fatal("CSM should be nil unless properly initialized by NewCrsceCompressor()")
	}
	if crsce.lsm != nil {
		t.Fatal("LSM should be nil unless properly initialized by NewCrsceCompressor()")
	}
	if crsce.vsm != nil {
		t.Fatal("VSM should be nil unless properly initialized by NewCrsceCompressor()")
	}
}
