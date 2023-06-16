package environment

import "testing"

// errorCheck compare two errors
func errorCheck(t *testing.T, name *string, actualError, expectedError error) {
	fatalError := func() {
		t.Fatalf("Error mismatch on %s\n"+
			"  actual:%v\n"+
			"expected:%v\n",
			*name, actualError, expectedError)
	}
	resolve := func(e error) string {
		if e == nil {
			return "_nil_"
		}
		return e.Error()
	}
	if resolve(actualError) != resolve(expectedError) {
		fatalError()
	}
}
