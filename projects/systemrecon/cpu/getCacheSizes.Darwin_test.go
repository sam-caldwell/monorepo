//go:build darwin
// +build darwin

package systemrecon

import (
	"errors"
	"github.com/sam-caldwell/go/v2/projects/convert"
	"strconv"
	"testing"

	runcommand "github.com/sam-caldwell/go/v2/projects/RunCommand"
)

func TestGetCacheSizes(t *testing.T) {
	tests := []struct {
		name        string
		mockOutput  string
		mockError   error
		expectedL1  int
		expectedL2  int
		expectedL3  int
		expectedErr error
	}{
		{
			name:        "valid output",
			mockOutput:  "32768",
			mockError:   nil,
			expectedL1:  32,
			expectedL2:  32,
			expectedL3:  32,
			expectedErr: nil,
		},
		{
			name:        "invalid output",
			mockOutput:  "invalid",
			mockError:   nil,
			expectedL1:  0,
			expectedL2:  0,
			expectedL3:  0,
			expectedErr: &strconv.NumError{Func: "Atoi", Num: "invalid", Err: strconv.ErrSyntax},
		},
		{
			name:        "execution error",
			mockOutput:  "",
			mockError:   errors.New("execution error"),
			expectedL1:  0,
			expectedL2:  0,
			expectedL3:  0,
			expectedErr: errors.New("execution error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := runcommand.MockCommandExecutor{
				Output: tt.mockOutput,
				Error:  tt.mockError,
			}

			l1, l2, l3, err := getCacheSizes(mock)
			if l1 != tt.expectedL1 || l2 != tt.expectedL2 || l3 != tt.expectedL3 {
				t.Errorf("Test (%s) failed:\n"+
					"\t got [%d, %d, %d]\n"+
					"\twant [%d, %d, %d]",
					tt.name, l1, l2, l3, tt.expectedL1, tt.expectedL2, tt.expectedL3)
			}
			if convert.ErrorToString(err) != convert.ErrorToString(tt.expectedErr) {
				t.Errorf("Test (%s) failed:\n"+
					"\t got '%v'\n"+
					"\twant '%v'",
					tt.name, err, tt.expectedErr)
			}
		})
	}
}
