package hijack

import (
	"reflect"
	"runtime"
	"testing"
)

func TestGetNameOfFunction(t *testing.T) {
	// Define a sample function
	sampleFunc := func() {
		// Function body
	}

	// Get the expected function name
	expectedName := runtime.FuncForPC(reflect.ValueOf(sampleFunc).Pointer()).Name()

	// Call the function under test
	actualName := getNameOfFunction(sampleFunc)

	// Compare the expected and actual names
	if actualName != expectedName {
		t.Errorf("Function name mismatch. Expected: %s, Got: %s", expectedName, actualName)
	}
}
