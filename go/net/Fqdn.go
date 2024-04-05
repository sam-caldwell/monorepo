package net

import (
	"fmt"
)

// Fqdn - Fully Qualified Domain Name
type Fqdn string

// FromString - Given a string pointer, convert it to an FQDN
func (o *Fqdn) FromString(v *string) (err error) {

	if IsValidAddress(*v) {
		*o = Fqdn(*v)
	} else {
		err = fmt.Errorf("FQDN cannot be empty")
	}
	return err
}
