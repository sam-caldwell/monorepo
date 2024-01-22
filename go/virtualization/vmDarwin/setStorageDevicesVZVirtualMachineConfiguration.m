/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract List of disk devices. Empty by default.
 @see VZVirtioBlockDeviceConfiguration
 */
void setStorageDevicesVZVirtualMachineConfiguration(void *config,
    void *storageDevices)
{
    if (@available(macOS 11, *)) {
        [(VZVirtualMachineConfiguration *)config setStorageDevices:[(NSMutableArray *)storageDevices copy]];
        return;
    }
    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
