package net

import "fmt"

// PortNumber - Network port number (0-65535)
type PortNumber uint16

func (o *PortNumber) FromInt(v uint) error {
	if v > 65535 {
		return fmt.Errorf("invalid number must be 0...65535")
	}
	*o = PortNumber(v)
	return nil
}
