package lfs

import "os/exec"

// detachLoopDevice detaches the given loop device.
func detachLoopDevice(loopDevice *string) {
	if _, err := exec.Command("losetup", "-d", *loopDevice).Output(); err != nil {
		panic(err)
	}
}
