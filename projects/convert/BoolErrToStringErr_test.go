package convert

import (
	"fmt"
	"testing"
)

func TestBoolErrToStringErr(t *testing.T) {
	testCases := []struct {
		name    string
		n       bool
		err     error
		wantStr string
		wantErr error
	}{
		{
			name:    "No error",
			n:       true,
			err:     nil,
			wantStr: "true",
			wantErr: nil,
		},
		{
			name:    "With error",
			n:       false,
			err:     fmt.Errorf("some error"),
			wantStr: "false",
			wantErr: fmt.Errorf("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotStr, gotErr := BoolErrToStringErr(tc.n, tc.err)

			if gotStr != tc.wantStr {
				t.Errorf("Expected string %q, got %q", tc.wantStr, gotStr)
			}

			if (gotErr == nil) != (tc.wantErr == nil) || (gotErr != nil && gotErr.Error() != tc.wantErr.Error()) {
				t.Errorf("Expected error %v, got %v", tc.wantErr, gotErr)
			}
		})
	}
}
