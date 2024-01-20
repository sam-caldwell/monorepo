package virtualization

import (
	"testing"
)

// TestVmName_Get - Test the VmName.Get() method
func TestVmName_Get(t *testing.T) {
	testFunc := func(nameValue string, expectedState string) {
		name := VmName(nameValue)

		if actual := name.Get(); actual != expectedState {
			t.Fatalf("invalid name should not be stored (got: %s) (expected:%s)", actual, expectedState)
		}
	}
	type TestData struct {
		actual   string
		expected string
	}
	testData := []TestData{
		{"vmName", "vmName"},
		{"vmName0", "vmName0"},
		{"vmName1", "vmName1"},
		{"vmName2", "vmName2"},
		{"vmName3", "vmName3"},
		{"vmName4", "vmName4"},
		{"vmName5", "vmName5"},
		{"vmName6", "vmName6"},
		{"vmName7", "vmName7"},
		{"vmName8", "vmName8"},
		{"vmName9", "vmName9"},
		{"vmName20", "vmName20"},
		{"vmName99", "vmName99"},
		{"vmName999", "vmName999"},
		{"vmName9999", "vmName9999"},
		{"vmName0valid", "vmName0valid"},
		{"vmName1valid", "vmName1valid"},
		{"vmName2valid", "vmName2valid"},
		{"vmName3valid", "vmName3valid"},
		{"vmName4valid", "vmName4valid"},
		{"vmName5valid", "vmName5valid"},
		{"vmName6valid", "vmName6valid"},
		{"vmName7valid", "vmName7valid"},
		{"vmName8valid", "vmName8valid"},
		{"vmName9valid", "vmName9valid"},
		{"vmName20valid", "vmName20valid"},
		{"vmName99valid", "vmName99valid"},
		{"vmName999valid", "vmName999valid"},
		{"vmName9999valid", "vmName9999valid"},
		{"vmName0valid0", "vmName0valid0"},
		{"vmName1valid1", "vmName1valid1"},
		{"vmName2valid2", "vmName2valid2"},
		{"vmName3valid3", "vmName3valid3"},
		{"vmName4valid4", "vmName4valid4"},
		{"vmName5valid5", "vmName5valid5"},
		{"vmName6valid6", "vmName6valid6"},
		{"vmName7valid7", "vmName7valid7"},
		{"vmName8valid8", "vmName8valid8"},
		{"vmName9valid9", "vmName9valid9"},
		{"vm-Name", "vm-Name"},
		{"vm.Name", "vm.Name"},
		{"vm_Name", "vm_Name"},
		{"vm,Name", "vm,Name"},
		{"vm>Name", "vm>Name"},
		{"vm<Name", "vm<Name"},
		{"vm[Name", "vm[Name"},
		{"vm]Name", "vm]Name"},
		{"vm|Name", "vm|Name"},
		{"vm\\Name", "vm\\Name"},
		{"vm/Name", "vm/Name"},
		{"vm?Name", "vm?Name"},
		{"vm`Name", "vm`Name"},
		{"vm~Name", "vm~Name"},
		{"vm!Name", "vm!Name"},
		{"vm@Name", "vm@Name"},
		{"vm#Name", "vm#Name"},
		{"vm$Name", "vm$Name"},
		{"vm%Name", "vm%Name"},
		{"vm^Name", "vm^Name"},
		{"vm&Name", "vm&Name"},
		{"vm*Name", "vm*Name"},
		{"vm(Name", "vm(Name"},
		{"vm)Name", "vm)Name"},
		{"vm+Name", "vm+Name"},
		{"vm=Name", "vm=Name"},
		{"vm:Name", "vm:Name"},
		{"vm;Name", "vm;Name"},
		{"vm'Name", "vm'Name"},
		{"vm\"Name", "vm\"Name"},
		{"0vmName", "0vmName"},
		{"1vmName", "1vmName"},
		{"2vmName", "2vmName"},
		{"3vmName", "3vmName"},
		{"4vmName", "4vmName"},
		{"5vmName", "5vmName"},
		{"6vmName", "6vmName"},
		{"7vmName", "7vmName"},
		{"8vmName", "8vmName"},
		{"9vmName", "9vmName"},
		{"60vmName", "60vmName"},
		{"999vmName", "999vmName"},
		{"9999vmName", "9999vmName"},
	}

	for _, row := range testData {
		testFunc(row.actual, row.expected)
	}
}
