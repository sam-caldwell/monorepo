package diskFormat

import "fmt"

func (format *DiskFormat) Set(f string) error {
	switch f {
	case "ext4":
		*format = Ext4
	case "swap":
		*format = Swap
	default:
		return fmt.Errorf("unknown, unrecognized or unsupported file format")
	}
	return nil
}
