package convert

import (
	"fmt"
	"testing"
)

func TestNumericErrToStringErr(t *testing.T) {
	//goland:noinspection GoRedundantConversion
	testCases := []struct {
		name    string
		n       any
		dt      string
		err     error
		wantStr string
		wantErr error
	}{
		{
			name:    "No error float64",
			n:       float64(1.23),
			dt:      "f64",
			err:     nil,
			wantStr: "1.23",
			wantErr: nil,
		},
		{
			name:    "With error float64",
			n:       float64(45.6),
			dt:      "f64",
			err:     fmt.Errorf("some error"),
			wantStr: "45.6",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "No error float32",
			n:       float32(1.24),
			dt:      "f32",
			err:     nil,
			wantStr: "1.24",
			wantErr: nil,
		},
		{
			name:    "With error float32",
			n:       float32(45.7),
			dt:      "f32",
			err:     fmt.Errorf("some error"),
			wantStr: "45.7",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "With error int",
			n:       int(45),
			dt:      "int",
			err:     fmt.Errorf("some error"),
			wantStr: "45",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "With error int8",
			n:       int8(45),
			dt:      "int8",
			err:     fmt.Errorf("some error"),
			wantStr: "45",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "With error int16",
			n:       int16(45),
			dt:      "int16",
			err:     fmt.Errorf("some error"),
			wantStr: "45",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "With error int32",
			n:       int32(45),
			dt:      "int32",
			err:     fmt.Errorf("some error"),
			wantStr: "45",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "With error int64",
			n:       int64(45),
			dt:      "int64",
			err:     fmt.Errorf("some error"),
			wantStr: "45",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "With error uint",
			n:       uint(45),
			dt:      "uint",
			err:     fmt.Errorf("some error"),
			wantStr: "45",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "With error uint8",
			n:       uint8(45),
			dt:      "uint8",
			err:     fmt.Errorf("some error"),
			wantStr: "45",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "With error uint16",
			n:       uint16(45),
			dt:      "uint16",
			err:     fmt.Errorf("some error"),
			wantStr: "45",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "With error uint32",
			n:       uint32(45),
			dt:      "uint32",
			err:     fmt.Errorf("some error"),
			wantStr: "45",
			wantErr: fmt.Errorf("some error"),
		},
		{
			name:    "With error uint64",
			n:       uint64(45),
			dt:      "uint64",
			err:     fmt.Errorf("some error"),
			wantStr: "45",
			wantErr: fmt.Errorf("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var gotStr string
			var gotErr error
			switch tc.dt {
			case "f64":
				fv := tc.n.(float64)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
			case "f32":
				fv := tc.n.(float32)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
			case "int":
				fv := tc.n.(int)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
			case "int8":
				fv := tc.n.(int8)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
			case "int16":
				fv := tc.n.(int16)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
			case "int32":
				fv := tc.n.(int32)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
			case "int64":
				fv := tc.n.(int64)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
			case "uint":
				fv := tc.n.(uint)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
			case "uint8":
				fv := tc.n.(uint8)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
			case "uint16":
				fv := tc.n.(uint16)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
			case "uint32":
				fv := tc.n.(uint32)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
			case "uint64":
				fv := tc.n.(uint64)
				gotStr, gotErr = NumericErrToStringErr(fv, tc.err)
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
