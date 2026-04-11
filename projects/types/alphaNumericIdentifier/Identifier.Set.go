package alphaNumericIdentifier

import (
	"fmt"
	"time"
)

const (
	defaultVmName = "vmInstance%d"
)

// Set - Set and validate the name identifier.
func (name *Identifier) Set(n string) error {
	if err := name.valid(&n); err != nil {
		*name = Identifier(fmt.Sprintf(defaultVmName, time.Now().UnixNano()))
		return err
	}
	*name = Identifier(n)
	return nil
}
