package virtualization

import (
	"fmt"
	"regexp"
	"testing"
)

// TestVmName_Set - Test the VmName.Set() method
func TestVmName_Set(t *testing.T) {
	testFunc := func(index int, nameValue string, expectedError error) {
		var name VmName

		//If the Set() method doesn't return the expected error, fail like a banana republic election.
		if err := name.Set(nameValue); (err != nil) && err.Error() != expectedError.Error() {
			t.Fatalf("%d: name (%s) expected error (%v) mismatch (%v)", index, nameValue, expectedError, err)
		}
		//If the Set() method returns the expected error...
		if expectedError == nil {
			//If expectedError is nil (no expected error) we expect our name to be properly stored.
			if string(name) != nameValue {
				t.Fatalf("%d: value mismatch (expected %s, got %s)", index, nameValue, string(name))
			}

		} else {
			pattern := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]+$")
			if !pattern.MatchString(string(name)) {
				t.Fatalf("%d: default name (%s) failed expected pattern", index, name)
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

	for index, row := range testData {
		testFunc(index, row.actual, row.err)
	}
}
