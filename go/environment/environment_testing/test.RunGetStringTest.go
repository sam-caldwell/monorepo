package environment_testing

import "testing"

// RunGetStringTest - Run the boolean (pointer) test
func RunGetStringTest(t *testing.T, name string, f TestStringFunc, value string, err error, expectedValue string) {
	Setup(t, name, value)
	if actualValue, actualError := f(name); actualError != nil {
		checkError(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		checkValue(t, &name, actualValue, expectedValue)
	}
}
