package net

import (
	"fmt"
	"strconv"
)

// PortNumber - Network port number (0-65535)
type PortNumber uint16

// FromInt - Get port number from integer
func (o *PortNumber) FromInt(v int) error {
	if v < 0 || v > 65535 {
		return fmt.Errorf("invalid number must be 0...65535")
	}
	*o = PortNumber(v)
	return nil
}

// FromString - Get port number from string
func (o *PortNumber) FromString(v string) error {
	n, err := strconv.Atoi(v)
	if err != nil {
		return err
	}
	return o.FromInt(n)
}
