package net

import "regexp"

// IsValidNetBIOSName - checks if the given string is a valid NETBIOS name
func IsValidNetBIOSName(s string) bool {
	// Ensure the length of the string is within the allowed range
	if len(s) < 1 || len(s) > 15 {
		return false
	}

	// Regular expression for valid NETBIOS name characters
	validNetBIOSName := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_\-\$]{0,14}$`)

	return validNetBIOSName.MatchString(s)
}
