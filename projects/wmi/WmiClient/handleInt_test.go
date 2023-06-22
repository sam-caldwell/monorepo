package wmiclient

import (
	"reflect"
	"testing"
)

func TestHandleInt(t *testing.T) {
	var intVal int = 10
	var int8Val int8 = 10
	var int16Val int16 = 10
	var int32Val int32 = 10
	var int64Val int64 = 10
	var uintVal uint = 10
	var uint8Val uint8 = 10
	var uint16Val uint16 = 10
	var uint32Val uint32 = 10
	var uint64Val uint64 = 10

	testCases := []struct {
		desc  string
		field reflect.Value
		value int64
	}{
		{"int", reflect.ValueOf(&intVal).Elem(), 20},
		{"int8", reflect.ValueOf(&int8Val).Elem(), 20},
		{"int16", reflect.ValueOf(&int16Val).Elem(), 20},
		{"int32", reflect.ValueOf(&int32Val).Elem(), 20},
		{"int64", reflect.ValueOf(&int64Val).Elem(), 20},
		{"uint", reflect.ValueOf(&uintVal).Elem(), 20},
		{"uint8", reflect.ValueOf(&uint8Val).Elem(), 20},
		{"uint16", reflect.ValueOf(&uint16Val).Elem(), 20},
		{"uint32", reflect.ValueOf(&uint32Val).Elem(), 20},
		{"uint64", reflect.ValueOf(&uint64Val).Elem(), 20},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := handleInt(tc.field, tc.value)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			switch tc.field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if int64(tc.field.Int()) != tc.value {
					t.Fatalf("expected %v, got %v", tc.value, tc.field.Int())
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if int64(tc.field.Uint()) != tc.value {
					t.Fatalf("expected %v, got %v", tc.value, tc.field.Uint())
				}
			}
		})
	}

	// test wrong kind case
	wrongKind := reflect.ValueOf("not an int")
	err := handleInt(wrongKind, 20)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
