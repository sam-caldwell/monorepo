//go:build windows
// +build windows

package wmiclient

import (
	"github.com/go-ole/go-ole"
	"reflect"
)

// handleSlice - Given a string slice, copy it from the OLE Property value to our internal state
func handleSlice(field reflect.Value, value *ole.SafeArrayConversion) (err error) {
	vector := map[reflect.Kind]func(s reflect.Value, v any){
		reflect.String: func(s reflect.Value, v any) { s.SetString(reflect.ValueOf(v).String()) },
		reflect.Int:    func(s reflect.Value, v any) { s.SetInt(reflect.ValueOf(v).Int()) },
		reflect.Uint:   func(s reflect.Value, v any) { s.SetUint(reflect.ValueOf(v).Uint()) },
	}
	if value != nil {
		arr := value.ToValueArray()
		fArr := reflect.MakeSlice(field.Type(), len(arr), len(arr))
		for i, v := range arr {
			s := fArr.Index(i)
			vector[field.Type().Kind()](s, v)
		}
		field.Set(fArr)
	}
	return err
}
