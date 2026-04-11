package alphaNumericIdentifier

import (
	"fmt"
	"regexp"
)

const (
	nameRegex      = "^[a-zA-Z][a-zA-Z0-9]+$"
	errInvalidName = "invalid name"
)

// valid - return error if the name is invalid or nil otherwise.
func (name *Identifier) valid(n *string) error {
	re := regexp.MustCompile(nameRegex)
	if re.MatchString(*n) {
		return nil
	}
	return fmt.Errorf(errInvalidName)
}
