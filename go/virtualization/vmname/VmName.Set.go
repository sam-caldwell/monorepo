package vmname

import (
	"fmt"
	"time"
)

const (
	defaultVmName = "vmInstance%d"
)

// Set - Set and validate the name identifier.
func (name *VmName) Set(n string) error {
	if err := name.valid(&n); err != nil {
		*name = VmName(fmt.Sprintf(defaultVmName, time.Now().UnixNano()))
		return err
	}
	*name = VmName(n)
	return nil
}
