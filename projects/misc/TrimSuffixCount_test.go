package misc

import "testing"

func TestTrimSuffixCount(t *testing.T) {
	tests := []struct {
		input    string
		suffix   string
		count    int
		expected string
	}{
		{"example_example_example", "_example", 0, "example_example_example"},
		{"example_example_example", "_example", 1, "example_example"},
		{"example_example_example", "_example", 2, "example"},
		{"test_test_test.txt", ".txt", 0, "test_test_test.txt"},
		{"test_test_test.txt", ".txt", 1, "test_test_test"},
		{"test_test_test.txt", ".txt", 2, "test_test_test"},
		{"no_suffix_here", "_suffix", 1, "no_suffix_here"},
		{"empty_string", "_string", 1, "empty"},
		{"empty_string", "_string", 2, "empty"},
		{"10.0.0.0", ".0", 1, "10.0.0"},
		{"10.0.0.0", ".0", 2, "10.0"},
		{"10.0.0.0", ".0", 3, "10"},
	}

	for _, test := range tests {
		result := TrimSuffixCount(test.input, test.suffix, test.count)

		if result != test.expected {
			t.Fatalf("For input %s with count %d, expected %s, but got %s", test.input, test.count, test.expected, result)
		}
	}
}
