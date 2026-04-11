package assert

import (
	"reflect"
	"testing"
)

// Equal - Print message and terminate test if a != b
func Equal(t *testing.T, lhs, rhs any, message string) {
	if !reflect.DeepEqual(lhs, rhs) {
		t.Fatal(message)
	}
}
