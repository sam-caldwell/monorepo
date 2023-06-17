package environment_testing

import (
	"testing"
)

// RunSanitizedIntTest - Run the SanitizedInt test
func RunSanitizedIntTest(t *testing.T, f TestIntSanityFunc, name, value string, err error, expectedValue, min, max int) {
	Setup(t, name, value)
	if actualValue, actualError := f(name, min, max); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}

// RunSanitizedIntpTest - Run the SanitizedInt (pointer) test
func RunSanitizedIntpTest(t *testing.T, f TestIntpSanityFunc, name, value string, err error, expectedValue, min, max int) {
	Setup(t, name, value)
	if actualValue, actualError := f(&name, min, max); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}
