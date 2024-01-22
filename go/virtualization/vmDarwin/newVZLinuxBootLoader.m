/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract Create a VZLinuxBootLoader with the Linux kernel passed as NSURL.
 @param kernelPath local file system path of the linux kernel.
*/
void *newVZLinuxBootLoader(const char *kernelPath)
{
    if (@available(macOS 11, *)) {
        NSString *kernelPathNSString = [NSString stringWithUTF8String:kernelPath];
        NSURL *kernelURL = [NSURL fileURLWithPath:kernelPathNSString];
        return [[VZLinuxBootLoader alloc] initWithKernelURL:kernelURL];
    }
    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
