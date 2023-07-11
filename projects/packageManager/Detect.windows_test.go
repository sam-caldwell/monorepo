//go:build windows
// +build windows

package packageManager

func TestDetect(t *testing.T) {
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
