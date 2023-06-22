package wmiclient

import (
	"reflect"
	"testing"
)

func TestHandleUint(t *testing.T) {
	var intVal int = 10
	var uintVal uint = 10

	testCases := []struct {
		desc  string
		field reflect.Value
		value uint64
	}{
		{"int", reflect.ValueOf(&intVal).Elem(), 20},
		{"uint", reflect.ValueOf(&uintVal).Elem(), 20},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := handleUint(tc.field, tc.value)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			switch tc.field.Kind() {
			case reflect.Int:
				if tc.field.Int() != int64(tc.value) {
					t.Fatalf("expected %v, got %v", tc.value, tc.field.Int())
				}
			case reflect.Uint:
				if tc.field.Uint() != tc.value {
					t.Fatalf("expected %v, got %v", tc.value, tc.field.Uint())
				}
			}
		})
	}

	// test wrong kind case
	wrongKind := reflect.ValueOf("string")
	err := handleUint(wrongKind, 20)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
