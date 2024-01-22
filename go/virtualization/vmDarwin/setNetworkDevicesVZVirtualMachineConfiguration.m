/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract List of network adapters. Empty by default.
 @see VZVirtioNetworkDeviceConfiguration
 */
void setNetworkDevicesVZVirtualMachineConfiguration(void *config,
    void *networkDevices)
{
    if (@available(macOS 11, *)) {
        [(VZVirtualMachineConfiguration *)config setNetworkDevices:[(NSMutableArray *)networkDevices copy]];
        return;
    }
    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
