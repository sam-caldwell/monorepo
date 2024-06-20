package lfs

import "github.com/sam-caldwell/monorepo/go/ansi"

// formatLoopDevice formats the given loop device to the specified filesystem format.
func formatLoopDevice(loopDevice *string, formatter formatFunction) error {

	out, err := formatter(loopDevice)

	if err == nil {
		ansi.Blue()
	} else {
		ansi.Red()
	}

	ansi.Printf("out: %v\n", out).Reset()

	return err
}
