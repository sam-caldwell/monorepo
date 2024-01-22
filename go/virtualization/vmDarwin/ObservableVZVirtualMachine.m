/*
 * Objective-C virtualization interfaces
 */
#import "virtualization_11.h"

@implementation ObservableVZVirtualMachine {
    Observer *_observer;
};
@implementation ObservableVZVirtualMachine {
    Observer *_observer;
};
- (instancetype)initWithConfiguration:(VZVirtualMachineConfiguration *)configuration
                                queue:(dispatch_queue_t)queue
                   statusUpdateHandle:(uintptr_t)statusUpdateHandle
{
    self = [super initWithConfiguration:configuration queue:queue];
    _observer = [[Observer alloc] init];
    [self addObserver:_observer
           forKeyPath:@"state"
              options:NSKeyValueObservingOptionNew
              context:(void *)statusUpdateHandle];
    return self;
}

- (void)dealloc
{
    [self removeObserver:_observer forKeyPath:@"state"];
    [_observer release];
    [super dealloc];
}
@end
