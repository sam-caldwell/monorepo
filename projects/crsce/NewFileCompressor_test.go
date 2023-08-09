package crsce

import "testing"

func TestNewCrsceCompressor(t *testing.T) {
	const (
		testData = `` +
			`0123456789` + // 010
			`0123456789` + // 020
			`0123456789` + // 030
			`0123456789` + // 040
			`0123456789` + // 050
			`0123456789` + // 060
			`0123456789` + // 070
			`0123456789` + // 080
			`0123456789` + // 090
			`0123456789` // 100
	)
	if len(testData) != 100 {
		t.Fatalf("testData size mismatch (actual: %d)", len(testData))
	}

	//crsce := NewCrsceCompressor(4)
}
