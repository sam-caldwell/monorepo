/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract Set the optional initial RAM disk path.
 @param bootLdrPtr VZLinuxBootLoader
 @param ramdiskPath The RAM disk is mapped into memory before booting the kernel.
 @link https://www.kernel.org/doc/html/latest/admin-guide/kernel-parameters.html
 */
void setInitialRamdiskURLVZLinuxBootLoader(void *bootLdrPtr, const char *ramdiskPath)
{
    if (@available(macOS 11, *)) {
        NSString *ramdiskPathNSString = [NSString stringWithUTF8String:ramdiskPath];
        NSURL *ramdiskURL = [NSURL fileURLWithPath:ramdiskPathNSString];
        [(VZLinuxBootLoader *)bootLdrPtr setInitialRamdiskURL:ramdiskURL];
        return;
    }
    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
