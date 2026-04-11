package lfs

import (
	"fmt"
)

// FormatPartition - Initialize a given file as a loop file to the given format.
//
// This function assumes `fileName` identifies a loop file on disk which can be
// formatted to the given format (e.g. ext4).
//
// Steps:
//   - verify format is supported
//   - map the loop device to the fileName
//   - defer detach the loop device
//   - format the file
func FormatPartition(fileName, format string) (err error) {
	var loopDevice string
	var formatter formatFunction
	fmt.Printf("Formatting partition %s to %s\n", fileName, format)

	if loopDevice, err = setupLoopDevice(fileName); err != nil {
		return fmt.Errorf("failed to set up loop device: %w", err)
	}
	defer detachLoopDevice(&loopDevice)
	if formatter, err = getFormatFunction(&format); err != nil {
		return err
	}
	if err = formatLoopDevice(&loopDevice, formatter); err != nil {
		return fmt.Errorf("failed to format loop device: %w", err)
	}
	return nil
}
