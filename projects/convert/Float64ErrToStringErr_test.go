package convert

import (
	"errors"
	"testing"
)

func TestFloat64ErrToStringErr(t *testing.T) {
	testCases := []struct {
		name    string
		n       float64
		err     error
		wantStr string
		wantErr error
	}{
		{
			name:    "No error",
			n:       1.23,
			err:     nil,
			wantStr: "1.230000",
			wantErr: nil,
		},
		{
			name:    "With error",
			n:       45.6,
			err:     errors.New("some error"),
			wantStr: "45.600000",
			wantErr: errors.New("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotStr, gotErr := Float64ErrToStringErr(tc.n, tc.err)

			if gotStr != tc.wantStr {
				t.Errorf("Expected string %q, got %q", tc.wantStr, gotStr)
			}

			if (gotErr == nil) != (tc.wantErr == nil) || (gotErr != nil && gotErr.Error() != tc.wantErr.Error()) {
				t.Errorf("Expected error %v, got %v", tc.wantErr, gotErr)
			}
		})
	}
}
