package valid

import (
	"fmt"
)

// IsValidHelpString - validate help string. return error state.
func IsValidHelpString(argument *string) error {
	if argument == nil {
		return fmt.Errorf(errEmptyOrNilObject)
	}
	return nil
}
