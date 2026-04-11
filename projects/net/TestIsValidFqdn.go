package net

import "regexp"

// IsValidDnsName - checks if the given string is a valid DNS name
func IsValidDnsName(input string) bool {
	// Regular expression pattern for valid DNS names
	dnsNamePattern := `^[a-zA-Z0-9]([a-zA-Z0-9\-]*[a-zA-Z0-9])?(?:\.[a-zA-Z0-9]([a-zA-Z0-9\-]*[a-zA-Z0-9])?)*$`

	// Compile the regular expression
	regex := regexp.MustCompile(dnsNamePattern)

	// Check if the input matches the regular expression
	return regex.MatchString(input)
}
