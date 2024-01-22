/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract List of serial ports. Empty by default.
 @see VZVirtioConsoleDeviceSerialPortConfiguration
 */
void setSerialPortsVZVirtualMachineConfiguration(void *config,
    void *serialPorts)
{
    if (@available(macOS 11, *)) {
        [(VZVirtualMachineConfiguration *)config setSerialPorts:[(NSMutableArray *)serialPorts copy]];
        return;
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
