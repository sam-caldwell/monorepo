package valid

import (
	"fmt"
	"regexp"
)

// IsNameArg - return nil error if *descriptor is a valid positional descriptor name
func IsNameArg(argument *string) error {
	if argument == nil {
		return fmt.Errorf(errEmptyOrNilObject)
	}
	if regexp.MustCompile(validArgRegex).MatchString(*argument) {
		return nil
	}
	return fmt.Errorf(errExpectedNameArg, *argument)
}
