package environment_testing

import (
	"testing"
)

// RunGetFloatTest - Run the Float test
func RunGetFloatTest(t *testing.T, name string, f TestFloatFunc, value string, err error, expectedValue float64) {
	Setup(t, name, value)
	if actualValue, actualError := f(name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}
