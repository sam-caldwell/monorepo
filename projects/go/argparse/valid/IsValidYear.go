package valid

import (
	"fmt"
	"regexp"
	"strings"
)

// IsValidYear - return nil error if input is a valid year.
func IsValidYear(year int) (string, error) {

	yearString := strings.TrimSpace(fmt.Sprintf("%d", year))

	if regexp.MustCompile(validYearRegex).MatchString(yearString) {
		return yearString, nil
	}
	return "", fmt.Errorf(errInvalidYear, year)
}
