/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract: Maximum amount of memory allowed for a virtual machine.
 @see VZVirtualMachineConfiguration.memorySize
 */
unsigned long long maximumAllowedMemorySizeVZVirtualMachineConfiguration()
{
    if (@available(macOS 11, *)) {
        return (unsigned long long)[VZVirtualMachineConfiguration maximumAllowedMemorySize];
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
