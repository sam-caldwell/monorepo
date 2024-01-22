/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract List of entropy devices. Empty by default.
 @see VZVirtioEntropyDeviceConfiguration
*/
void setEntropyDevicesVZVirtualMachineConfiguration(void *config,
    void *entropyDevices)
{
    if (@available(macOS 11, *)) {
        [(VZVirtualMachineConfiguration *)config setEntropyDevices:[(NSMutableArray *)entropyDevices copy]];
        return;
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
