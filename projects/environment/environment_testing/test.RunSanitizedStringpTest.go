package environment_testing

import "testing"

// RunSanitizedStringpTest - Run the SanitizedStringp (pointer) test
func RunSanitizedStringpTest(t *testing.T, f TestStringpSanityFunc, err error, name, value, pattern, expectedValue string) {
	Setup(t, name, value)
	if actualValue, actualError := f(&name, &pattern); actualError != nil {
		checkError(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		checkValue(t, &name, actualValue, expectedValue)
	}
}
