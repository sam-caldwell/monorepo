package logger

import (
	"regexp"
)

// ValidAppName - Return boolean if application name is valid
//
//	(c) 2023 Sam Caldwell.  MIT License
func ValidAppName(value *string) bool {
	pattern := regexp.MustCompile(`[a-zA-Z0-9_][a-zA-Z0-9_\.\-]{0,255}`)
	return pattern.MatchString(*value)
}
