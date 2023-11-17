package simple

import (
	"github.com/sam-caldwell/monorepo/go/misc"
	"reflect"
	"testing"
)

func TestSet_GetType(t *testing.T) {
	var set Set
	set.Init()
	if set.GetType() != reflect.UnsafePointer {
		t.Fatal("Expected initial type to be reflect.UnsafePointer")
	}
	//goland:noinspection GoRedundantConversion
	testData := map[any]reflect.Kind{
		int(1):       reflect.Int,
		int8(2):      reflect.Int8,
		int16(3):     reflect.Int16,
		int32(4):     reflect.Int32,
		int64(5):     reflect.Int64,
		1.0:          reflect.Float64,
		float32(4.1): reflect.Float32,
		true:         reflect.Bool,
		false:        reflect.Bool,
		uint(6):      reflect.Uint,
		uint8(7):     reflect.Uint8,
		uint16(8):    reflect.Uint16,
		uint32(9):    reflect.Uint32,
		uint64(10):   reflect.Uint64,
		"test":       reflect.String,
	}

	for value, typ := range testData {
		if actual := reflect.TypeOf(value).Kind(); actual != typ {
			t.Fatalf("Test error.  Got '%s' Expected '%s'", actual, typ)
		}
		set.data[value] = misc.NullObjectStruct{}
		if actual := set.GetType(); actual != typ {
			t.Fatalf("Given '%v' expected '%s' Got: '%s'", value, typ.String(), actual.String())
		}
		delete(set.data, value)
	}
}
