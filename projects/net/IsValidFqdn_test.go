package net

import "testing"

func TestIsValidFqdn(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"example.com", true},                // Valid DNS name
		{"subdomain.example.com", true},      // Valid DNS subdomain
		{"host.subdomain.example.com", true}, // Valid DNS subdomain
		{"example.com.", true},               // Valid DNS name with trailing ..
		{"example.com..", false},             // Valid DNS name with trailing ...
		{"example.com...", false},            // Valid DNS name with trailing dot
		{"www.example.com", true},            // Valid DNS name
		{"-example.com", false},              // Invalid DNS name (starts with '-')
		{"example-.com", false},              // Invalid DNS name (ends with '-')
		{"ex--ample.com", false},             // Invalid DNS name (consecutive dashes)
		{"ex..ample.com", false},             //Invalid DNS name (contains ..)
		{"example.com-", false},              // Invalid DNS name (ends with '-')
		{"ex@mple.com", false},               // Invalid DNS name (contains special character)
		{"", false},                          // Invalid DNS name (empty string)
	}

	for _, testCase := range testCases {
		actual := IsValidAddress(testCase.input)
		if actual != testCase.expected {
			t.Fatalf("ERR: For input '%s', expected %t but got %t", testCase.input, testCase.expected, actual)
		}
	}
}
