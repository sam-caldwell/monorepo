package nettools

import (
	"testing"
)

func TestGetNetworkAddress(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"192.168.1.1/24", "192.168.1.0", false},
		{"10.0.0.1/16", "10.0.0.0", false},
		{"invalidCIDR", "", true},
	}

	for _, test := range tests {
		result, err := GetNetworkAddress(test.input)

		if test.wantErr && err == nil {
			t.Fatalf("Expected an error for input %s, but got none.", test.input)
		}

		if !test.wantErr && err != nil {
			t.Fatalf("Unexpected error for input %s: %s", test.input, err)
		}

		if result != test.expected {
			t.Fatalf("For input %s, expected %s, but got %s", test.input, test.expected, result)
		}
	}
}
