package environment_testing

import "testing"

// RunGetIntpTest - Run the Int (pointer) test
func RunGetIntpTest(t *testing.T, name string, f TestIntpFunc, value string, err error, expectedValue int) {
	Setup(t, name, value)
	if actualValue, actualError := f(&name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}
