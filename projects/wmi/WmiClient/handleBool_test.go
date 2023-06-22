package wmiclient

import (
	"reflect"
	"testing"
)

func TestHandleBool(t *testing.T) {
	// Test cases
	cases := []struct {
		name  string
		field reflect.Value
		value bool
		err   bool
	}{
		{
			name:  "set true to bool",
			field: reflect.ValueOf(new(bool)).Elem(),
			value: true,
			err:   false,
		},
		{
			name:  "set false to bool",
			field: reflect.ValueOf(new(bool)).Elem(),
			value: false,
			err:   false,
		},
		{
			name:  "set true to non-bool",
			field: reflect.ValueOf(new(int)).Elem(),
			value: true,
			err:   true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := handleBool(tc.field, tc.value)
			if (err != nil) != tc.err {
				t.Fatalf("expected error: %v, got: %v", tc.err, err != nil)
			}
			if !tc.err && tc.field.Bool() != tc.value {
				t.Fatalf("expected value: %v, got: %v", tc.value, tc.field.Bool())
			}
		})
	}
}
