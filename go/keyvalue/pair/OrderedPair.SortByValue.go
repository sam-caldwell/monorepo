package pair

import (
	"github.com/sam-caldwell/monorepo/go/misc"
	"reflect"
	"sort"
	"unsafe"
)

func (o *OrderedPair) SortByValue() {
	compareSlice := func(valueI, valueJ reflect.Value) bool {
		// Handle []byte or other slice types
		// You can define your own comparison logic here
		return false
	}
	compareStruct := func(valueI, valueJ reflect.Value) bool {
		// Compare binary equivalence of structs
		ptrI := unsafe.Pointer(valueI.Pointer())
		ptrJ := unsafe.Pointer(valueJ.Pointer())
		size := valueI.Type().Size()
		return misc.CompareByteSlice((*[1 << 30]byte)(ptrI)[:size], (*[1 << 30]byte)(ptrJ)[:size]) == misc.BytesLess
	}
	sort.Slice(o.data, func(i, j int) bool {
		valueI := reflect.ValueOf(o.data[i].Value)
		valueJ := reflect.ValueOf(o.data[j].Value)

		switch valueI.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return valueI.Int() > valueJ.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return valueI.Uint() > valueJ.Uint()
		case reflect.Float32, reflect.Float64:
			return valueI.Float() > valueJ.Float()
		case reflect.String:
			return valueI.String() > valueJ.String()
		case reflect.Slice:
			return compareSlice(valueI, valueJ)
		case reflect.Struct:
			return compareStruct(valueI, valueJ)
		default:
			// Handle unsupported types or custom comparisons
			// You can define your own comparison logic here
			return false
		}
	})
}
