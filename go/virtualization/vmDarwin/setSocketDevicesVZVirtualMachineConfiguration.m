/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract List of socket devices. Empty by default.
 @see VZVirtioSocketDeviceConfiguration
 */
void setSocketDevicesVZVirtualMachineConfiguration(void *config,
    void *socketDevices)
{
    if (@available(macOS 11, *)) {
        [(VZVirtualMachineConfiguration *)config setSocketDevices:[(NSMutableArray *)socketDevices copy]];
        return;
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
