//go:build darwin
// +build darwin

package packageManager

import (
	"errors"
	"reflect"
	"testing"
)

func TestDetect(t *testing.T) {
	pkg, err := Detect()
	expectedPkg := brew
	expectedErr := error(nil)

	if !errors.Is(expectedErr, err) {
		t.Fatalf("error mismatch")
	}

	if reflect.ValueOf(pkg).Pointer() != reflect.ValueOf(expectedPkg).Pointer() {
		t.Fatalf("package mismatch")
	}
}
