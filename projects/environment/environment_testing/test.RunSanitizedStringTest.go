package environment_testing

import "testing"

// RunSanitizedStringTest - Run the SanitizedString test
func RunSanitizedStringTest(t *testing.T, f TestStringSanityFunc, err error, name, value, pattern, expectedValue string) {
	Setup(t, name, value)
	if actualValue, actualError := f(name, pattern); actualError != nil {
		checkError(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		checkValue(t, &name, actualValue, expectedValue)
	}
}
