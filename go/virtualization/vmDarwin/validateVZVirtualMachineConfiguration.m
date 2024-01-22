/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract Validate the configuration.
 @param config  Virtual machine configuration.
 @param error If not nil, assigned with the validation error if the validation failed.
 @return true if the configuration is valid.
 */
bool validateVZVirtualMachineConfiguration(void *config, void **error)
{
    if (@available(macOS 11, *)) {
        return (bool)[(VZVirtualMachineConfiguration *)config
            validateWithError:(NSError *_Nullable *_Nullable)error];
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
