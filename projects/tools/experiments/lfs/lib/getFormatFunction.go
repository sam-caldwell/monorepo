package lfs

import (
	"fmt"
	"os/exec"
)

type formatFunction func(f *string) (out string, err error)

// getFormatFunction - return the formatter function
func getFormatFunction(format *string) (formatFunction, error) {

	switch *format {
	case "ext4":

		return func(f *string) (out string, err error) {
			var raw []byte
			raw, err = exec.Command(`mkfs.ext4`, *f).Output()
			return string(raw), err
		}, nil

	case "swap":

		return func(f *string) (out string, err error) {
			var raw []byte
			raw, err = exec.Command(`mkswap`, *f).Output()
			return string(raw), err
		}, nil

	default:
		return nil, fmt.Errorf("invalid format")
	}

}
