//go:build linux
// +build linux

package packageManager

import (
	"testing"
)

func TestAptGet(t *testing.T) {
	// Mock DependencyDescriptor
	pkg := DependencyDescriptor{
		Name:           "test-package",
		Detail:         "test-package",
		SkipIfDetected: false,
	}

	// Mock runcommand.ShellExecute to return expected output and nil error
	//runcommand.ShellExecute = func(cmd string) (string, error) {
	//	return "Package installed successfully", nil
	//}

	// Mock systemrecon.HasExecutable to return exit code 0
	//systemrecon.HasExecutable = func(name string) (int, error) {
	//	return 0, nil
	//}

	// Call aptGet function
	output, err := aptGet(pkg)

	// Verify the output and error
	expectedOutput := "Package installed successfully"
	if output != expectedOutput {
		t.Errorf("Unexpected output. Expected: %s, Got: %s", expectedOutput, output)
	}
	if err != nil {
		t.Errorf("Unexpected error. Got: %v", err)
	}
}
