//go:build windows
// +build windows

package wmiclient

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/misc"
	"reflect"
)

const (
	/*
		VtNull (or VT_NULL) is a value type defined in the Microsoft Component Object Model (COM) specification.
		It is used to denote a property that does not have a value, similar to how you might use null in other
		programming contexts.

		In COM, VtNull is one of many variant types (VT_*) that can be used to specify the type of VARIANT, a
		special COM type that can contain different kinds of data.

		A VARIANT has a vt field (variant type) that indicates what kind of data it currently holds. The VtNull
		variant type indicates that the VARIANT holds no data.
	*/
	VtNull = 0x01
)

type NULL struct{}

// loadEntity loads a SWbemObject into a struct pointer.
func (client *Client) loadEntity(dst interface{}, src *ole.IDispatch) (err error) {
	v := reflect.ValueOf(dst).Elem()
	for fieldPos := 0; fieldPos < v.NumField(); fieldPos++ {
		/*
		 * To make the defer used below safer, we have wrapped this
		 * in a function.
		 */
		err := func() (err error) {
			/*
			 * Collect what we know about the field at the current position (fieldPos)
			 * and bail as early as possible.
			 */
			thisFieldValue := v.Field(fieldPos)
			originalFieldValue := thisFieldValue
			thisFieldName := v.Type().Field(fieldPos).Name
			if !misc.IsStringCapitalized(thisFieldName) {
				return nil
			}

			//If CanSet() is false, the field is read-only
			if !thisFieldValue.CanSet() {
				return fmt.Errorf(errors.ReadOnly+errors.Details, thisFieldValue)
			}

			// Fetch the COM OLE property
			oleProperty, err := oleutil.GetProperty(src, thisFieldName)
			if err != nil {
				if client.AllowMissingFields {
					return nil
				}
				return fmt.Errorf(errors.MissingField+errors.Details, thisFieldName)
			}
			defer errors.Swallow(oleProperty.Clear())
			/*
			 * Handle the situation where we are passing a pointer.
			 */
			isPointer := thisFieldValue.Kind() == reflect.Ptr

			// if we have a pointer, make sure nothing is null.
			if isPointer && !(client.PtrNil && isOleVariantPropertyNull(oleProperty.VT)) {
				ptr := reflect.New(thisFieldValue.Type().Elem())
				thisFieldValue.Set(ptr)
				thisFieldValue = thisFieldValue.Elem()
				//ToDo: should we just bail here for another pass?
			}
			if isOleVariantPropertyNull(oleProperty.VT) {
				return nil //Bail!
			}
			/*
			 * Evaluate the property value and cast the property value
			 * into our field
			 */
			switch thisPropertyValue := oleProperty.Value().(type) {
			case int8, int16, int32, int64, int:
				return handleInt(thisFieldValue, reflect.ValueOf(thisPropertyValue).Int())
			case uint8, uint16, uint32, uint64:
				return handleUint(thisFieldValue, reflect.ValueOf(thisPropertyValue).Uint())
			case string:
				return handleString(thisFieldValue, thisPropertyValue)
			case bool:
				return handleBool(thisFieldValue, thisPropertyValue)
			case float32:
				return handleFloat32(thisFieldValue, thisPropertyValue)
			case float64:
				return handleFloat64(thisFieldValue, thisPropertyValue)
			}
			/* We don't know what this is...more investigation...
			 * Let's look at this from the other side...
			 */
			if thisFieldValue.Kind() == reflect.Slice {
				supportedSliceTypes := map[reflect.Kind]NULL{
					reflect.String: NULL{}, reflect.Uint8: NULL{}, reflect.Uint16: NULL{}, reflect.Uint32: NULL{},
					reflect.Uint64: NULL{}, reflect.Uint: NULL{}, reflect.Int8: NULL{}, reflect.Int16: NULL{},
					reflect.Int32: NULL{}, reflect.Int64: NULL{}, reflect.Int: NULL{},
				}
				if _, found := supportedSliceTypes[thisFieldValue.Type().Elem().Kind()]; found {
					return handleSlice(thisFieldValue, oleProperty.ToArray())
				}
				err = fmt.Errorf(errors.TypeMismatch+errors.Details+errors.Details,
					thisFieldValue.Kind(), thisFieldValue.Type().Elem().Kind())
			} else {
				thisPropertyValueType := reflect.TypeOf(oleProperty.Value())
				if (thisPropertyValueType == nil) && (isPointer || client.NonePtrZero) {
					return StoreZeroValueBasedOnTypeIfNull(isPointer, client, originalFieldValue)
				}
				err = fmt.Errorf("unsupported type (%v)", thisPropertyValueType)
			}
			return nil
		}() //end of function

		if err != nil {
			if err.Error() == "break" {
				break
			}
			return err
		}

	}
	return err
}
