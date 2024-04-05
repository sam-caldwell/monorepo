package types

import (
	"fmt"
	"testing"
)

func TestApiKey_FromString(t *testing.T) {
	var apiKey ApiKey
	test := func(inp string, expected error) {
		t.Run(fmt.Sprintf("Input: %s", inp), func(t *testing.T) {
			err := apiKey.FromString(&inp)
			if err != nil {
				if expected == nil {
					t.Fatalf("unexpected error (%v) on '%s'", err, inp)
				}
				if err.Error() != expected.Error() {
					t.Errorf("Expected error: %v, got: %v", expected, err)
				}
			}
			if expected != nil && err == nil {
				t.Errorf("Expected error: %v, got nil", expected)
			}
		})
	}

	testCases := map[string]error{
		"MGUzZTc1MjM0YWJjNjhmNDM3OGE4NmIzZjRiMzJhMTk4YmEzMDE4NDViMGNkNmU1": nil,                          // Valid base64 string with padding
		"MGUzZTc1MjM0YWJjNjhmNDM3OGE4NmIzZjRiMzJhMTk4YmEzMDE4NDViMGNkNmU":  fmt.Errorf("invalid apiKey"), // short
		"validBase64String":        fmt.Errorf("invalid apiKey"), // Missing padding
		"invalidCharacter$Input==": fmt.Errorf("invalid apiKey"), // Invalid character
		"shortString":              fmt.Errorf("invalid apiKey"), // Short string
		"MGUzZTc1MjM0YWJjNjhmNDM3OGE4NmIzZjRiMzJhMTk4YmEzMDE4NDViMGNkNmU19": fmt.Errorf("invalid apiKey"), // too long
		"tooLongBase64String==tooLongBase64String==":                        fmt.Errorf("invalid apiKey"), // invalid (padding)
	}

	// Iterate over test cases
	for key, value := range testCases {
		test(key, value)
	}
}
