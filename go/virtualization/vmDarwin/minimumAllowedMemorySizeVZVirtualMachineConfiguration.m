/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract: Minimum amount of memory required by virtual machines.
 @see VZVirtualMachineConfiguration.memorySize
 */
unsigned long long minimumAllowedMemorySizeVZVirtualMachineConfiguration()
{
    if (@available(macOS 11, *)) {
        return (unsigned long long)[VZVirtualMachineConfiguration minimumAllowedMemorySize];
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
