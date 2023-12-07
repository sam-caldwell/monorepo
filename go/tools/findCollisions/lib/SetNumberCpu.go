package findCollision

import (
	"log"
	"runtime"
)

// SetNumberCpu - Set the number of workers/processes
func SetNumberCpu() int {
	//
	// The number of CPUs will determine the number
	// of workers we will distribute the problem across.
	//
	numCpu := runtime.NumCPU()

	log.Printf("Number CPU: %d", numCpu)

	if runtime.NumCPU() < 8 {
		panic("at least 8 CPU cores is required")
	}
	if runtime.NumCPU() > 255 {
		panic("more than 255 CPUs not supported at this time")
	}

	runtime.GOMAXPROCS(numCpu)

	return numCpu

}
