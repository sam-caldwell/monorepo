package wmiclient

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestHandleString(t *testing.T) {
	var stringVal string = "old_string"
	var intVal int = 10
	var uintVal uint = 10
	var timeVal time.Time = time.Now()

	testCases := []struct {
		desc  string
		field reflect.Value
		value string
	}{
		{"string", reflect.ValueOf(&stringVal).Elem(), "new_string"},
		{"int", reflect.ValueOf(&intVal).Elem(), "20"},
		{"uint", reflect.ValueOf(&uintVal).Elem(), "20"},
		{"time", reflect.ValueOf(&timeVal).Elem(), time.Now().Add(time.Hour).Format("20060102150405.000000-0700")},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := handleString(tc.field, tc.value)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
				return
			}
			switch tc.field.Kind() {
			case reflect.String:
				if tc.field.String() != tc.value {
					t.Errorf("expected %v, got %v", tc.value, tc.field.String())
				}
			case reflect.Int:
				expectedInt, err := strconv.Atoi(tc.value)
				if err != nil {
					t.Errorf("error converting string to int: %v", err)
					return
				}
				if int(tc.field.Int()) != expectedInt {
					t.Errorf("expected %v, got %v", expectedInt, tc.field.Int())
				}
			case reflect.Uint:
				expectedUint, err := strconv.ParseUint(tc.value, 10, 64)
				if err != nil {
					t.Errorf("error converting string to uint: %v", err)
					return
				}
				if tc.field.Uint() != expectedUint {
					t.Errorf("expected %v, got %v", expectedUint, tc.field.Uint())
				}
			case reflect.Struct:
				parsedTime, _ := time.Parse("20060102150405.000000-0700", tc.value)
				if tc.field.Interface().(time.Time) != parsedTime {
					t.Errorf("expected %v, got %v", parsedTime, tc.field.Interface().(time.Time))
				}
			}
		})
	}

	// test wrong kind case
	wrongKind := reflect.ValueOf(1.23)
	err := handleString(wrongKind, "20")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
