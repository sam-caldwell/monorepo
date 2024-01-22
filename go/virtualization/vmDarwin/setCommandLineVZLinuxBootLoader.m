/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract Set bool loader command-line parameters.
 @param bootLdrPtr VZLinuxBootLoader
 @param commandLine - cli kernel boot parameters.
 @link https://www.kernel.org/doc/html/latest/admin-guide/kernel-parameters.html
 */
void setCommandLineVZLinuxBootLoader(void *bootLdrPtr, const char *commandLine)
{
    if (@available(macOS 11, *)) {
        NSString *commandLineNSString = [NSString stringWithUTF8String:commandLine];
        [(VZLinuxBootLoader *)bootLdrPtr setCommandLine:commandLineNSString];
        return;
    }
    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}

