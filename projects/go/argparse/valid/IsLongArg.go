package valid

import (
	"fmt"
	"regexp"
)

// IsLongArg - return nil error if *descriptor is a valid long descriptor (--{string})
func IsLongArg(argument *string) error {
	if argument == nil {
		return fmt.Errorf(errEmptyOrNilObject)
	}
	if regexp.MustCompile(validLongArgRegex).MatchString(*argument) {
		return nil
	}
	return fmt.Errorf(errExpectedLongArg, *argument)
}
