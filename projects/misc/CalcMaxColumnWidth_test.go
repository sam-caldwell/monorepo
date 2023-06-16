package misc

import "testing"

func TestCalcMaxColumnWidth(t *testing.T) {
	data := []string{
		"",
		"a",
		"boo",
		"test",
		"test_me",
		"test_now",
		"test__now",
		"what_size_Here",
		"test__now",
		"what_size_Here_what_size_Here",
		"what_size_Here",
	}
	output := []int{
		0,
		0,
		1,
		3,
		4,
		7,
		8,
		9,
		14,
		14,
		29,
		29,
	}

	for n := range data {
		thisList := data[:n]
		if max := CalcMaxColumnWidth(&thisList); max != output[n] {
			t.Fatalf("max not correct [%d] (%d):%v", n, output, thisList)
		}
	}

	if w := CalcMaxColumnWidth(nil); w != 0 {
		t.Fail()
	}
}
