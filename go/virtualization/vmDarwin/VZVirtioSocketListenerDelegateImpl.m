/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

@implementation VZVirtioSocketListenerDelegateImpl {
    uintptr_t _cgoHandle;
}

- (instancetype)initWithHandle:(uintptr_t)cgoHandle
{
    self = [super init];
    _cgoHandle = cgoHandle;
    return self;
}

- (BOOL)listener:(VZVirtioSocketListener *)listener shouldAcceptNewConnection:(VZVirtioSocketConnection *)connection fromSocketDevice:(VZVirtioSocketDevice *)socketDevice;
{
    return (BOOL)shouldAcceptNewConnectionHandler(_cgoHandle, connection, socketDevice);
}
@end
