package simpleArgs

import (
	"os"
	"testing"
)

func TestGetCommand(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		args        []string
		helpText    string
		expectedCmd string
	}{
		{
			args:        []string{"program", "command"},
			helpText:    "Usage: program [command]",
			expectedCmd: "command",
		},
	}

	// Iterate over test cases
	for _, testCase := range testCases {
		// Set the command-line arguments
		os.Args = testCase.args

		// Call the GetCommand function
		actualCmd := GetCommand(testCase.helpText)

		// Verify the returned command
		if actualCmd != testCase.expectedCmd {
			t.Fatalf("Expected command: %s, but got: %s", testCase.expectedCmd, actualCmd)
		}
	}
}
