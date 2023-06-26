//go:build windows
// +build windows

package systemrecon

import (
	"errors"
	runcommand "github.com/sam-caldwell/go/v2/projects/runcommand"
	"testing"
)

func TestGetCacheSizes(t *testing.T) {
	testCases := []struct {
		name          string
		level         int
		mockOutput    string
		mockError     error
		expectedSize  int
		expectedError error
	}{
		{
			name:          "Valid L1 Cache",
			level:         1,
			mockOutput:    "BlockSize=64\nNumberOfBlocks=1024\n",
			mockError:     nil,
			expectedSize:  64 * 1024,
			expectedError: nil,
		},
		{
			name:          "Invalid Level",
			level:         4,
			mockOutput:    "",
			mockError:     nil,
			expectedSize:  invalidCacheSz,
			expectedError: errors.New("Invalid cache level: 4"),
		},
		{
			name:          "Command Execution Error",
			level:         1,
			mockOutput:    "",
			mockError:     errors.New("Command Execution Error"),
			expectedSize:  invalidCacheSz,
			expectedError: errors.New("Command Execution Error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockExecutor := runcommand.MockCommandExecutor{
				Output: tc.mockOutput,
				Error:  tc.mockError,
			}
			size, err := getCacheSizes(mockExecutor, tc.level)
			if size != tc.expectedSize || (err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error()) {
				t.Errorf("Expected size: %v, error: %v, but got size: %v, error: %v",
					tc.expectedSize, tc.expectedError, size, err)
			}
		})
	}
}
