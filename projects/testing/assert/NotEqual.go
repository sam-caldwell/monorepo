package assert

import (
	"reflect"
	"testing"
)

// NotEqual - Print message and terminate test if a == b
func NotEqual(t *testing.T, lhs, rhs any, message string) {
	if reflect.DeepEqual(lhs, rhs) {
		t.Fatal(message)
	}
}
