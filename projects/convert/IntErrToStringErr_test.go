package convert

import (
	"errors"
	"testing"
)

func TestIntErrToStringErr(t *testing.T) {
	testCases := []struct {
		name    string
		n       int
		err     error
		wantStr string
		wantErr error
	}{
		{
			name:    "No error",
			n:       123,
			err:     nil,
			wantStr: "123",
			wantErr: nil,
		},
		{
			name:    "With error",
			n:       456,
			err:     errors.New("some error"),
			wantStr: "456",
			wantErr: errors.New("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotStr, gotErr := IntErrToStringErr(tc.n, tc.err)

			if gotStr != tc.wantStr {
				t.Errorf("Expected string %q, got %q", tc.wantStr, gotStr)
			}

			if (gotErr == nil) != (tc.wantErr == nil) || (gotErr != nil && gotErr.Error() != tc.wantErr.Error()) {
				t.Errorf("Expected error %v, got %v", tc.wantErr, gotErr)
			}
		})
	}
}
