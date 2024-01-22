/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

/*!
 @abstract Initialize the VZFileHandleSerialPortAttachment from file descriptors.
 @param readFileDescriptor File descriptor for reading from the file.
 @param writeFileDescriptor File descriptor for writing to the file.
 @discussion
    Each file descriptor must a valid.
*/
void *newVZFileHandleSerialPortAttachment(int readFileDescriptor, int writeFileDescriptor)
{
    if (@available(macOS 11, *)) {
        VZFileHandleSerialPortAttachment *ret;
        @autoreleasepool {
            NSFileHandle *fileHandleForReading = [[NSFileHandle alloc] initWithFileDescriptor:readFileDescriptor];
            NSFileHandle *fileHandleForWriting = [[NSFileHandle alloc] initWithFileDescriptor:writeFileDescriptor];
            ret = [[VZFileHandleSerialPortAttachment alloc]
                initWithFileHandleForReading:fileHandleForReading
                        fileHandleForWriting:fileHandleForWriting];
        }
        return ret;
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
