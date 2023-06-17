package environment_testing

import "testing"

// checkValue compare two values
func checkValue(t *testing.T, name *string, actual, expected any) {
	if actual != expected {
		t.Fatalf("value mismatch on %s\n"+
			"  actual:%v\n"+
			"expected:%v\n",
			*name,
			actual,
			expected)
	}
}
