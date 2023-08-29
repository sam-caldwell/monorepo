package environment_testing

import "testing"

// RunSanitizedIntpTest - Run the SanitizedInt (pointer) test
func RunSanitizedIntpTest(t *testing.T, f TestIntpSanityFunc, name, value string, err error, expectedValue, min, max int) {
	Setup(t, name, value)
	if actualValue, actualError := f(&name, min, max); actualError != nil {
		checkError(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		checkValue(t, &name, actualValue, expectedValue)
	}
}
