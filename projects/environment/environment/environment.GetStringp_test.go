package environment

import (
	"os"
	"reflect"
	"testing"
)

func TestGetStringp(t *testing.T) {
	setup := func(n string, v string) {
		if err := os.Setenv(n, v); err != nil {
			t.Fatal(err)
		}
	}

	test := func(name string, value string, err error, expectedValue string) {
		setup(name, value)
		if actualValue, actualError := GetStringp(&name); actualError != nil {
			if reflect.DeepEqual(actualError, err) {
				t.Fatalf("Error mismatch on %s\n"+
					"  actual:%v\n"+
					"expected:%v\n", name, actualError, err)
			}
		} else if actualValue != expectedValue {
			t.Fatalf("value mismatch on %s\n"+
				"  actual:%v\n"+
				"expected:%v\n", name, actualValue, expectedValue)
		}
	}

	test("testBool0", "true", nil, "true")
	test("testBool1", "false", nil, "false")
	test("testBool2", "some string", nil, "some string")
	test("testBool3", "", nil, "")
}
