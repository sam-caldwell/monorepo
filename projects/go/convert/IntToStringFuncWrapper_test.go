package convert

import (
	"fmt"
	"testing"
)

func TestIntToString(t *testing.T) {
	tests := []struct {
		name string
		fn   func() (int, error)
		want string
		err  error
	}{
		{
			name: "With valid int",
			fn: func() (int, error) {
				return 123, nil
			},
			want: "123",
			err:  nil,
		},
		{
			name: "With error",
			fn: func() (int, error) {
				return 0, fmt.Errorf("test error")
			},
			want: "0",
			err:  fmt.Errorf("test error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToStringFuncWrapper(tt.fn)
			if got != tt.want {
				t.Errorf("IntToStringFuncWrapper() = %v, want %v", got, tt.want)
			}
			if (err != nil && tt.err == nil) || (err == nil && tt.err != nil) || (err != nil && tt.err != nil && err.Error() != tt.err.Error()) {
				t.Errorf("IntToStringFuncWrapper() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
