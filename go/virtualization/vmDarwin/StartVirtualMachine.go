package vmDarwin

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Virtualization
#import <Foundation/Foundation.h>
#import <Virtualization/Virtualization.h>

BOOL startVirtualMachine(VZVirtualMachine* virtualMachine) {
    NSError* error = nil;
    if (![virtualMachine startWithError:&error]) {
        NSLog(@"Error starting virtual machine: %@", error);
        return NO;
    }
    return YES;
}
*/
import "C"

// StartVirtualMachine starts the given virtual machine.
func StartVirtualMachine(virtualMachine *C.VZVirtualMachine) bool {
	return C.startVirtualMachine(virtualMachine) != 0
}
