/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract: Maximum number of CPUs for a virtual machine.
 @see VZVirtualMachineConfiguration.CPUCount
 */
unsigned int maximumAllowedCPUCountVZVirtualMachineConfiguration()
{
    if (@available(macOS 11, *)) {
        return (unsigned int)[VZVirtualMachineConfiguration maximumAllowedCPUCount];
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
