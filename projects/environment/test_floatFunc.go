package environment

import (
	"testing"
)

// runFloatTest - run the Float test
func runFloatTest(t *testing.T, name string, f testFloatFunc, value string, err error, expectedValue float64) {
	setup(t, name, value)
	if actualValue, actualError := f(name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}

// runFloatpTest - run the Float (pointer) test
func runFloatpTest(t *testing.T, name string, f testFloatpFunc, value string, err error, expectedValue float64) {
	setup(t, name, value)
	if actualValue, actualError := f(&name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}

// runSanitizedFloatTest - run the SanitizedFloat test
func runSanitizedFloatTest(t *testing.T, f testFloatSanityFunc, name, value string, err error, expectedValue, min, max float64) {
	setup(t, name, value)
	if actualValue, actualError := f(name, min, max); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}

// runSanitizedFloatTest - run the SanitizedFloat (pointer) test
func runSanitizedFloatpTest(t *testing.T, f testFloatpSanityFunc, name, value string, err error, expectedValue, min, max float64) {
	setup(t, name, value)
	if actualValue, actualError := f(&name, min, max); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}
