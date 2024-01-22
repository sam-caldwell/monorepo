/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract Return the list of network devices configurations for this VZVirtualMachineConfiguration. Return an empty array if no network device configuration is set.
 */
void *networkDevicesVZVirtualMachineConfiguration(void *config)
{
    if (@available(macOS 11, *)) {
        return [(VZVirtualMachineConfiguration *)config networkDevices]; // NSArray<VZSocketDeviceConfiguration *>
    }
    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
