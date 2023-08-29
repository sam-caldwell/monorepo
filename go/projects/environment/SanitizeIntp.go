package environment

import "fmt"

// SanitizeIntp - return sanitized environment variable (sanitized)
func SanitizeIntp(name *string, min, max int) (int, error) {
	value, err := GetIntp(name)
	if err != nil {
		return 0, err
	}
	if (value < min) || (value > max) {
		return 0, fmt.Errorf(errEnvBoundCheck)
	}
	return value, err
}
