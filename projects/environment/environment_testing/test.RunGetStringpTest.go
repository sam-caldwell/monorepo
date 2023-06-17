package environment_testing

import "testing"

// RunGetStringpTest - Run the boolean test
func RunGetStringpTest(t *testing.T, name string, f TestStringpFunc, value string, err error, expectedValue string) {
	Setup(t, name, value)
	if actualValue, actualError := f(&name); actualError != nil {
		errorCheck(t, &name, actualError, err)
	} else if actualValue != expectedValue {
		t.Fatalf("value mismatch on %s\n"+
			"  actual:'%v'\n"+
			"expected:'%v'\n", name, actualValue, expectedValue)
	}
}
