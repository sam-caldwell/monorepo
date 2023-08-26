package ordered

import "testing"

func TestSet_Has(t *testing.T) {
	var set Set
	test := func(value any, expected bool) {
		if set.Has(nil) != expected {
			t.Fail()
		}
	}
	testData := map[any]bool{
		nil:      false,
		0:        false,
		1:        false,
		0.0:      false,
		1.0:      false,
		true:     false,
		false:    false,
		"string": false,
	}
	for value, expected := range testData {
		test(value, expected)
	}
	_ = set.Add("myString")
	if !set.Has("myString") {
		t.Fail()
	}
}
