package wmiclient

import (
	"reflect"
	"testing"
)

func TestStrToUint64(t *testing.T) {
	var myUint uint64
	myUintVal := reflect.ValueOf(&myUint).Elem()
	err := strToUint64(myUintVal, "123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if myUint != 123 {
		t.Fatalf("expected %v, got %v", 123, myUint)
	}

	// Test with a non-uint value.
	myString := "not a uint"
	err = strToUint64(myUintVal, myString)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}
