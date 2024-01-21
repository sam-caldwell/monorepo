package vmDarwin

/*

#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Virtualization
#import <Foundation/Foundation.h>
#import <Virtualization/Virtualization.h>

VZVirtualMachineConfiguration* createVirtualMachineConfiguration(int cpuCount, long long memory) {
    VZVirtualMachineConfiguration* configuration = [[VZVirtualMachineConfiguration alloc] init];
    configuration.CPUCount = cpuCount;
    configuration.memorySize = memory;
    // Set up other configurations as needed
    return configuration;
}

*/

import "C"

// CreateVirtualMachineConfiguration creates a virtual machine configuration with the specified CPU count and memory size.
func CreateVirtualMachineConfiguration(cpuCount int, memory int64) *C.VZVirtualMachineConfiguration {
	return C.createVirtualMachineConfiguration(C.int(cpuCount), C.longlong(memory))
}
