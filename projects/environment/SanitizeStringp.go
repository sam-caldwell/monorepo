package environment

import (
	"fmt"
	"regexp"
)

// SanitizeStringp  - return the string env var (sanitized by regex pattern)
func SanitizeStringp(name, pattern *string) (string, error) {
	value, err := GetStringp(name)
	if err != nil {
		return value, err
	}
	if regexp.MustCompile(*pattern).MatchString(value) {
		return value, nil
	}
	return defaultStringValue, fmt.Errorf(errPatternCheckFailed)
}
