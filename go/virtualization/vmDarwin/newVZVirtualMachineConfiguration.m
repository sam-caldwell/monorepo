/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract Create a new Virtual machine configuration.
 @param bootLoader Boot loader used when the virtual machine starts.

 @param CPUCount Number of CPUs on the target machine
 @discussion
    Value range:
        VZVirtualMachineConfiguration.minimumAllowedCPUCount ... VZVirtualMachineConfiguration.maximumAllowedCPUCount.

 @see VZVirtualMachineConfiguration.minimumAllowedCPUCount
 @see VZVirtualMachineConfiguration.maximumAllowedCPUCount

 @param memorySize Virtual machine memory size in bytes.
 @discussion
    The memory size must be a multiple of a 1 megabyte (1024 * 1024 bytes) between VZVirtualMachineConfiguration.minimumAllowedMemorySize
    and VZVirtualMachineConfiguration.maximumAllowedMemorySize.

    The memorySize represents the total physical memory seen by a guest OS running in the virtual machine.
    Not all memory is allocated on start, the virtual machine allocates memory on demand.
 @see VZVirtualMachineConfiguration.minimumAllowedMemorySize
 @see VZVirtualMachineConfiguration.maximumAllowedMemorySize
 */
void *newVZVirtualMachineConfiguration(void *bootLoaderPtr,
    unsigned int CPUCount,
    unsigned long long memorySize)
{
    if (@available(macOS 11, *)) {
        VZVirtualMachineConfiguration *config = [[VZVirtualMachineConfiguration alloc] init];
        [config setBootLoader:(VZLinuxBootLoader *)bootLoaderPtr];
        [config setCPUCount:(NSUInteger)CPUCount];
        [config setMemorySize:memorySize];
        return config;
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
