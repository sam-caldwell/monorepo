package environment_testing

import (
	"testing"
)

// RunSanitizedStringTest - Run the SanitizedString test
func RunSanitizedStringTest(t *testing.T, f TestStringSanityFunc, err error, name, value, pattern, expectedValue string) {
	Setup(t, name, value)
	if actualValue, actualError := f(name, pattern); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}

// RunSanitizedStringpTest - Run the SanitizedStringp (pointer) test
func RunSanitizedStringpTest(t *testing.T, f TestStringpSanityFunc, err error, name, value, pattern, expectedValue string) {
	Setup(t, name, value)
	if actualValue, actualError := f(&name, &pattern); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}
