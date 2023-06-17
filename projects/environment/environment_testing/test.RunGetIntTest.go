package environment_testing

import "testing"

// RunGetIntTest - Run the Int test
func RunGetIntTest(t *testing.T, name string, f TestIntFunc, value string, err error, expectedValue int) {
	Setup(t, name, value)
	if actualValue, actualError := f(name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		t.Fatalf("value mismatch on %s\n"+
			"  actual:%v\n"+
			"expected:%v\n", name, actualValue, expectedValue)
	}
}
