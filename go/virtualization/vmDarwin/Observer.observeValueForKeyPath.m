/*
 * Objective-C virtualization interfaces
 */
 #import "virtualization_11.h"

@implementation Observer
- (void)observeValueForKeyPath:(NSString *)key ofObject:(id)object change:(NSDictionary *)change context:(void *)context;
{
    if ([key isEqualToString:@"state"]) {
        int state = (int)[change[NSKeyValueChangeNewKey] integerValue];
        changeStateOnObserver(state, (uintptr_t)context);
    }
}
@end
