package net

// IsValidAddress - test whether the input string is a valid IP address or FQDN
func IsValidAddress(s string) bool {
	return IsValidIpAddress(s) || IsValidFqdn(s)
}
