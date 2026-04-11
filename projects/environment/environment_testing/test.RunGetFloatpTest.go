package environment_testing

import "testing"

// RunGetFloatpTest - Run the Float (pointer) test
func RunGetFloatpTest(t *testing.T, name string, f TestFloatpFunc, value string, err error, expectedValue float64) {
	Setup(t, name, value)
	if actualValue, actualError := f(&name); actualError != nil {
		checkError(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		checkValue(t, &name, actualValue, expectedValue)
	}
}
