//go:build windows
// +build windows

package wmiclient

// coinitService coinitializes WMI service. If no error is returned, a cleanup function
// is returned which must be executed (usually deferred) to clean up allocated resources.
func (client *Client) coinitService(connectServerArgs ...interface{}) (*ole.IDispatch, func(), error) {
	var unknown *ole.IUnknown
	var wmi *ole.IDispatch
	var serviceRaw *ole.VARIANT

	// be sure teardown happens in the reverse
	// order from that which they were created
	deferFn := func() {
		if serviceRaw != nil {
			serviceRaw.Clear()
		}
		if wmi != nil {
			wmi.Release()
		}
		if unknown != nil {
			unknown.Release()
		}
		ole.CoUninitialize()
	}

	// if we error'ed here, clean up immediately
	var err error
	defer func() {
		if err != nil {
			deferFn()
		}
	}()

	err = ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	if err != nil {
		oleCode := err.(*ole.OleError).Code()
		if oleCode != ole.S_OK && oleCode != WindowsManagementInstrumentation.S_FALSE {
			return nil, nil, err
		}
	}

	unknown, err = oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		return nil, nil, err
	} else if unknown == nil {
		return nil, nil, WindowsManagementInstrumentation.ErrNilCreateObject
	}

	wmi, err = unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, nil, err
	}

	// service is a SWbemServices
	serviceRaw, err = oleutil.CallMethod(wmi, "ConnectServer", connectServerArgs...)
	if err != nil {
		return nil, nil, err
	}

	return serviceRaw.ToIDispatch(), deferFn, nil
}
