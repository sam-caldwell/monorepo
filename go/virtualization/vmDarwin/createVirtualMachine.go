package vmDarwin

import "C"
import "unsafe"

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Virtualization
#import <Foundation/Foundation.h>
#import <Virtualization/Virtualization.h>

VZVirtualMachine* createVirtualMachine(VZVirtualMachineConfiguration* configuration, NSString* name, NSString* identifier) {
    NSError* error = nil;
    VZVirtualMachine* virtualMachine = [[VZVirtualMachine alloc] initWithConfiguration:configuration name:name identifier:identifier error:&error];
    if (error != nil) {
        NSLog(@"Error creating virtual machine: %@", error);
        return nil;
    }
    return virtualMachine;
}
*/

// CreateVirtualMachine creates a virtual machine with the given configuration, name, and identifier.
func CreateVirtualMachine(configuration *C.VZVirtualMachineConfiguration, name string, identifier string) *C.VZVirtualMachine {

	nameStr := C.CString(name)
	defer C.free(unsafe.Pointer(nameStr))

	identifierStr := C.CString(identifier)
	defer C.free(unsafe.Pointer(identifierStr))

	return C.createVirtualMachine(configuration, C.CString(name), C.CString(identifier))
}
