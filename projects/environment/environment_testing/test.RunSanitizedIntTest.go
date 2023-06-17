package environment_testing

import "testing"

// RunSanitizedIntTest - Run the SanitizedInt test
func RunSanitizedIntTest(t *testing.T, f TestIntSanityFunc, name, value string, err error, expectedValue, min, max int) {
	Setup(t, name, value)
	if actualValue, actualError := f(name, min, max); actualError != nil {
		checkError(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		checkValue(t, &name, actualValue, expectedValue)
	}
}
