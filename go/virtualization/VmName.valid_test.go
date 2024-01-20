package virtualization

import (
	"fmt"
	"testing"
)

func TestVmName_valid(t *testing.T) {

	testFunc := func(index int, expected string, expectedError error) {
		var name VmName

		if err := name.valid(&expected); (err != nil) && err.Error() != expectedError.Error() {
			t.Fatalf("name (%s) expected error state (%v) mismatch (%v)", expected, err, expectedError)
		}

		if expectedError == nil {
			if string(name) != "" {
				t.Fatalf("we should not have set any value since no error was returned.")
			}
		} else {
			if string(name) != "" {
				t.Fatalf("%d: given (%s) default name (%s) failed expected pattern", index, expected, name)
			}
		}
	}
	type TestData struct {
		actual string
		err    error
	}
	testData := []TestData{
		{"vmName", nil},
		{"vmName0", nil},
		{"vmName1", nil},
		{"vmName2", nil},
		{"vmName3", nil},
		{"vmName4", nil},
		{"vmName5", nil},
		{"vmName6", nil},
		{"vmName7", nil},
		{"vmName8", nil},
		{"vmName9", nil},
		{"vmName20", nil},
		{"vmName99", nil},
		{"vmName999", nil},
		{"vmName9999", nil},
		{"vmName0valid", nil},
		{"vmName1valid", nil},
		{"vmName2valid", nil},
		{"vmName3valid", nil},
		{"vmName4valid", nil},
		{"vmName5valid", nil},
		{"vmName6valid", nil},
		{"vmName7valid", nil},
		{"vmName8valid", nil},
		{"vmName9valid", nil},
		{"vmName20valid", nil},
		{"vmName99valid", nil},
		{"vmName999valid", nil},
		{"vmName9999valid", nil},
		{"vmName0valid0", nil},
		{"vmName1valid1", nil},
		{"vmName2valid2", nil},
		{"vmName3valid3", nil},
		{"vmName4valid4", nil},
		{"vmName5valid5", nil},
		{"vmName6valid6", nil},
		{"vmName7valid7", nil},
		{"vmName8valid8", nil},
		{"vmName9valid9", nil},
		{"vm-Name", fmt.Errorf(errInvalidName)},
		{"vm.Name", fmt.Errorf(errInvalidName)},
		{"vm_Name", fmt.Errorf(errInvalidName)},
		{"vm,Name", fmt.Errorf(errInvalidName)},
		{"vm>Name", fmt.Errorf(errInvalidName)},
		{"vm<Name", fmt.Errorf(errInvalidName)},
		{"vm[Name", fmt.Errorf(errInvalidName)},
		{"vm]Name", fmt.Errorf(errInvalidName)},
		{"vm|Name", fmt.Errorf(errInvalidName)},
		{"vm\\Name", fmt.Errorf(errInvalidName)},
		{"vm/Name", fmt.Errorf(errInvalidName)},
		{"vm?Name", fmt.Errorf(errInvalidName)},
		{"vm`Name", fmt.Errorf(errInvalidName)},
		{"vm~Name", fmt.Errorf(errInvalidName)},
		{"vm!Name", fmt.Errorf(errInvalidName)},
		{"vm@Name", fmt.Errorf(errInvalidName)},
		{"vm#Name", fmt.Errorf(errInvalidName)},
		{"vm$Name", fmt.Errorf(errInvalidName)},
		{"vm%Name", fmt.Errorf(errInvalidName)},
		{"vm^Name", fmt.Errorf(errInvalidName)},
		{"vm&Name", fmt.Errorf(errInvalidName)},
		{"vm*Name", fmt.Errorf(errInvalidName)},
		{"vm(Name", fmt.Errorf(errInvalidName)},
		{"vm)Name", fmt.Errorf(errInvalidName)},
		{"vm+Name", fmt.Errorf(errInvalidName)},
		{"vm=Name", fmt.Errorf(errInvalidName)},
		{"vm:Name", fmt.Errorf(errInvalidName)},
		{"vm;Name", fmt.Errorf(errInvalidName)},
		{"vm'Name", fmt.Errorf(errInvalidName)},
		{"vm\"Name", fmt.Errorf(errInvalidName)},
		{"0vmName", fmt.Errorf(errInvalidName)},
		{"1vmName", fmt.Errorf(errInvalidName)},
		{"2vmName", fmt.Errorf(errInvalidName)},
		{"3vmName", fmt.Errorf(errInvalidName)},
		{"4vmName", fmt.Errorf(errInvalidName)},
		{"5vmName", fmt.Errorf(errInvalidName)},
		{"6vmName", fmt.Errorf(errInvalidName)},
		{"7vmName", fmt.Errorf(errInvalidName)},
		{"8vmName", fmt.Errorf(errInvalidName)},
		{"9vmName", fmt.Errorf(errInvalidName)},
		{"60vmName", fmt.Errorf(errInvalidName)},
		{"999vmName", fmt.Errorf(errInvalidName)},
		{"9999vmName", fmt.Errorf(errInvalidName)},
	}

	for n, row := range testData {
		testFunc(n, row.actual, row.err)
	}
}
