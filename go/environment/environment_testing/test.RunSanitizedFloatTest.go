package environment_testing

import "testing"

// RunSanitizedFloatTest - Run the SanitizedFloat test
func RunSanitizedFloatTest(t *testing.T, f TestFloatSanityFunc, name, value string, err error, expectedValue, min, max float64) {
	Setup(t, name, value)
	if actualValue, actualError := f(name, min, max); actualError != nil {
		checkError(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		checkValue(t, &name, actualValue, expectedValue)
	}
}
