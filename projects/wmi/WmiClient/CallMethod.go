//go:build windows
// +build windows

package wmiclient

import (
	"fmt"
	"github.com/go-ole/go-ole/oleutil"
)

// CallMethod calls a WMI method named methodName on an instance
// of the class named className. It passes in the arguments given
// in params. Use connectServerArgs to customize the machine and
// namespace; by default, the local machine and default namespace
// are used. See
// https://docs.microsoft.com/en-us/windows/desktop/WmiSdk/swbemlocator-connectserver
// for details.
func (client *Client) CallMethod(connectServerArgs []interface{}, className, methodName string, params []interface{}) (int32, error) {
	service, cleanup, err := client.coinitService(connectServerArgs...)
	if err != nil {
		return 0, fmt.Errorf("coinit: %v", err)
	}
	defer cleanup()

	// Get class
	classRaw, err := oleutil.CallMethod(service, "Get", className)
	if err != nil {
		return 0, fmt.Errorf("CallMethod Get class %s: %v", className, err)
	}
	class := classRaw.ToIDispatch()
	defer classRaw.Clear()

	// Run method
	resultRaw, err := oleutil.CallMethod(class, methodName, params...)
	if err != nil {
		return 0, fmt.Errorf("CallMethod %s.%s: %v", className, methodName, err)
	}
	resultInt, ok := resultRaw.Value().(int32)
	if !ok {
		return 0, fmt.Errorf("return value was not an int32: %v (%T)", resultRaw, resultRaw)
	}

	return resultInt, nil
}
