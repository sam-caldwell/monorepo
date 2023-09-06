//go:build !windows && !linux && !darwin
// +build !windows,!linux,!darwin

package packageManager

func TestDetect(t *testing.T) {
	f, err := Detect()
	if f != nil {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}
