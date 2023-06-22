package wmiclient

import (
	"reflect"
	"testing"
)

func TestHandleFloat32(t *testing.T) {
	var f float32 = 1.234

	// test success case
	successCase := reflect.ValueOf(&f).Elem()
	err := handleFloat32(successCase, f)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if f != 1.234 {
		t.Fatalf("expected 1.234, got %v", f)
	}

	// test wrong kind case
	wrongKind := reflect.ValueOf("not a float32")
	err = handleFloat32(wrongKind, f)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
