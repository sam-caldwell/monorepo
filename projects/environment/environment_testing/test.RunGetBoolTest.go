package environment_testing

import "testing"

// RunGetBoolTest - Run the boolean test
func RunGetBoolTest(t *testing.T, name string, f TestBoolFunc, value string, err error, expectedValue bool) {
	Setup(t, name, value)
	if actualValue, actualError := f(name); actualError != nil {
		checkError(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		checkValue(t, &name, actualValue, expectedValue)
	}
}
