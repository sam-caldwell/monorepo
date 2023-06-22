//go:build windows
// +build windows

package wmiclient

import (
	"github.com/go-ole/go-ole"
	"reflect"
)

// handleStringSlice - Given a string slice, copy it from the OLE Property value to our internal state
func handleStringSlice(field reflect.Value, value *ole.SafeArrayConversion) (err error) {
	if value != nil {
		arr := value.ToValueArray()
		fArr := reflect.MakeSlice(field.Type(), len(arr), len(arr))
		for i, v := range arr {
			s := fArr.Index(i)
			s.SetString(v.(string))
		}
		field.Set(fArr)
	}
	return err
}

func handleIntSlice(field reflect.Value, value *ole.SafeArrayConversion) (err error) {
	if value != nil {
		arr := value.ToValueArray()
		fArr := reflect.MakeSlice(field.Type(), len(arr), len(arr))
		for i, v := range arr {
			s := fArr.Index(i)
			s.SetInt(reflect.ValueOf(v).Int())
		}
		field.Set(fArr)
	}
	return err
}
func handleUintSlice(field reflect.Value, value *ole.SafeArrayConversion) (err error) {
	if value != nil {
		arr := value.ToValueArray()
		fArr := reflect.MakeSlice(field.Type(), len(arr), len(arr))
		for i, v := range arr {
			s := fArr.Index(i)
			s.SetUint(reflect.ValueOf(v).Uint())
		}
		field.Set(fArr)
	}
	return err
}