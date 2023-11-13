package convert

import (
	"fmt"
	"testing"
)

func TestNumericErrToStringErr(t *testing.T) {
	testCases := []struct {
		name    string
		n       any
		dt      string
		err     error
		wantStr string
		wantErr error
	}{
		{
			name:    "No error 64",
			n:       float64(1.23),
			dt:      "64",
			err:     nil,
			wantStr: "1.23",
			wantErr: nil,
		},
		{
			name:    "With error 64",
			n:       float64(45.6),
			dt:      "64",
			err:     fmt.Errorf("some error"),
			wantStr: "45.6",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "No error 32",
			n:       float32(1.24),
			dt:      "32",
			err:     nil,
			wantStr: "1.24",
			wantErr: nil,
		},
		{
			name:    "With error 32",
			n:       float32(45.7),
			dt:      "32",
			err:     fmt.Errorf("some error"),
			wantStr: "45.7",
			wantErr: fmt.Errorf("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var gotStr string
			var gotErr error
			switch tc.dt {
			case "64":
				fv := tc.n.(float64)
				gotStr, gotErr = FloatErrToStringErr(fv, tc.err)
			case "32":
				fv := tc.n.(float32)
				gotStr, gotErr = FloatErrToStringErr(fv, tc.err)
			default:
				t.Fatal("invalid typ")
			}
			if gotStr != tc.wantStr {
				t.Errorf("Expected string %q, got %q", tc.wantStr, gotStr)
			}

			if (gotErr == nil) != (tc.wantErr == nil) || (gotErr != nil && gotErr.Error() != tc.wantErr.Error()) {
				t.Errorf("Expected error %v, got %v", tc.wantErr, gotErr)
			}
		})
	}
}
