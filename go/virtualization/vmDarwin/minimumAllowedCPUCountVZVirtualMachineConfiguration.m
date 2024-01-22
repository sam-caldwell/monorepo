/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract: Minimum number of CPUs for a virtual machine.
 @see VZVirtualMachineConfiguration.CPUCount
 */
unsigned int minimumAllowedCPUCountVZVirtualMachineConfiguration()
{
    if (@available(macOS 11, *)) {
        return (unsigned int)[VZVirtualMachineConfiguration minimumAllowedCPUCount];
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
