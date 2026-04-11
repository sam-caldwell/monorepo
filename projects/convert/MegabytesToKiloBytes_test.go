package convert

import (
	"testing"
)

func TestMegabytesToKiloBytes(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			name: "1 megabyte to kilobytes",
			arg:  1,
			want: 1024,
		},
		{
			name: "0 megabytes to kilobytes",
			arg:  0,
			want: 0,
		},
		{
			name: "negative megabytes to kilobytes",
			arg:  -1,
			want: -1024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MegabytesToKiloBytes(tt.arg); got != tt.want {
				t.Fatalf("MegabytesToKiloBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
