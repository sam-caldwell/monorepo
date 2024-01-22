/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract Return the list of socket devices configurations for this VZVirtualMachineConfiguration. Return an empty array if no socket device configuration is set.
 */
void *socketDevicesVZVirtualMachineConfiguration(void *config)
{
    if (@available(macOS 11, *)) {
        return [(VZVirtualMachineConfiguration *)config socketDevices]; // NSArray<VZSocketDeviceConfiguration *>
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
