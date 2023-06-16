package environment

import (
	"testing"
)

// runIntTest - run the Int test
func runIntTest(t *testing.T, name string, f testIntFunc, value string, err error, expectedValue int) {
	setup(t, name, value)
	if actualValue, actualError := f(name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		t.Fatalf("value mismatch on %s\n"+
			"  actual:%v\n"+
			"expected:%v\n", name, actualValue, expectedValue)
	}
}

// runIntpTest - run the Int (pointer) test
func runIntpTest(t *testing.T, name string, f testIntpFunc, value string, err error, expectedValue int) {
	setup(t, name, value)
	if actualValue, actualError := f(&name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}

// runSanitizedIntTest - run the SanitizedInt test
func runSanitizedIntTest(t *testing.T, f testIntSanityFunc, name, value string, err error, expectedValue, min, max int) {
	setup(t, name, value)
	if actualValue, actualError := f(name, min, max); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}

// runSanitizedIntpTest - run the SanitizedInt (pointer) test
func runSanitizedIntpTest(t *testing.T, f testIntpSanityFunc, name, value string, err error, expectedValue, min, max int) {
	setup(t, name, value)
	if actualValue, actualError := f(&name, min, max); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		valueCheck(t, &name, actualValue, expectedValue)
	}
}
