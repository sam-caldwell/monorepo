package convert

import (
	"testing"
)

func TestUnsafeAtoI(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			name: "Valid integer string",
			arg:  "123",
			want: 123,
		},
		{
			name: "Invalid integer string",
			arg:  "abc",
			want: 0,
		},
		{
			name: "Empty string",
			arg:  "",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnsafeAtoI(tt.arg); got != tt.want {
				t.Errorf("UnsafeAtoI() = %v, want %v", got, tt.want)
			}
		})
	}
}
