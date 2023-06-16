package environment

import (
	"testing"
)

// runStringTest - run the boolean test
func runStringTest(t *testing.T, name string, f testStringFunc, value string, err error, expectedValue string) {
	setup(t, name, value)
	if actualValue, actualError := f(name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		t.Fatalf("value mismatch on %s\n"+
			"  actual:'%v'\n"+
			"expected:'%v'\n", name, actualValue, expectedValue)
	}
}

// runStringpTest - run the boolean (pointer) test
func runStringpTest(t *testing.T, name string, f testStringpFunc, value string, err error, expectedValue string) {
	setup(t, name, value)
	if actualValue, actualError := f(&name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}

// runSanitizedStringTest - run the SanitizedString test
func runSanitizedStringTest(t *testing.T, f testStringSanityFunc, err error, name, value, pattern, expectedValue string) {
	setup(t, name, value)
	if actualValue, actualError := f(name, pattern); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}

// runSanitizedStringpTest - run the SanitizedStringp (pointer) test
func runSanitizedStringpTest(t *testing.T, f testStringpSanityFunc, err error, name, value, pattern, expectedValue string) {
	setup(t, name, value)
	if actualValue, actualError := f(&name, &pattern); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}
