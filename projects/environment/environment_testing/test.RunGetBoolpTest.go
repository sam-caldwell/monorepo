package environment_testing

import (
	"testing"
)

// RunGetBoolpTest - Run the boolean (pointer) test
func RunGetBoolpTest(t *testing.T, name string, f TestBoolpFunc, value string, err error, expectedValue bool) {
	Setup(t, name, value)
	if actualValue, actualError := f(&name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}
