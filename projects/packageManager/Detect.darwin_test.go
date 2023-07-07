//go:build darwin
// +build darwin

package packageManager

import (
	"reflect"
	"testing"
)

func TestDetect(t *testing.T) {
	// Test case 1: Supported package manager detected
	pkg, err := Detect()
	expectedPkg := brew
	expectedErr := error(nil)

	if err != expectedErr {
		t.Fatalf("error mismatch")
	}

	if reflect.ValueOf(pkg).Pointer() != reflect.ValueOf(expectedPkg).Pointer() {
		t.Fatalf("package mismatch")
	}
}
