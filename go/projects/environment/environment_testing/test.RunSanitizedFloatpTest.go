package environment_testing

import "testing"

// RunSanitizedFloatpTest - Run the SanitizedFloat (pointer) test
func RunSanitizedFloatpTest(t *testing.T, f TestFloatpSanityFunc, name, value string, err error, expectedValue, min, max float64) {
	Setup(t, name, value)
	if actualValue, actualError := f(&name, min, max); actualError != nil {
		checkError(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		checkValue(t, &name, actualValue, expectedValue)
	}
}
