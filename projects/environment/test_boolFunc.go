package environment

import (
	"testing"
)

// runBoolTest - run the boolean test
func runBoolTest(t *testing.T, name string, f testBoolFunc, value string, err error, expectedValue bool) {
	setup(t, name, value)
	if actualValue, actualError := f(name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}

// runBoolpTest - run the boolean (pointer) test
func runBoolpTest(t *testing.T, name string, f testBoolpFunc, value string, err error, expectedValue bool) {
	setup(t, name, value)
	if actualValue, actualError := f(&name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}
