/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract List of memory balloon devices. Empty by default.
 @see VZVirtioTraditionalMemoryBalloonDeviceConfiguration
*/
void setMemoryBalloonDevicesVZVirtualMachineConfiguration(void *config,
    void *memoryBalloonDevices)
{
    if (@available(macOS 11, *)) {
        [(VZVirtualMachineConfiguration *)config setMemoryBalloonDevices:[(NSMutableArray *)memoryBalloonDevices copy]];
        return;
    }
    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
