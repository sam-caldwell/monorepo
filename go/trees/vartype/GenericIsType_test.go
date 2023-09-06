package vartype

import (
	"reflect"
	"testing"
)

func TestGeneric_IsType(t *testing.T) {
	var o Generic

	test := func(name string, data any, typ reflect.Kind) {
		if o.data = data; !o.IsType(typ) {
			t.Fatalf("Fail on %s", name)
		}
	}

	test("bool true", true, reflect.Bool)
	test("bool false", false, reflect.Bool)

	test("float32 -pi", float32(-3.141529), reflect.Float32)
	test("float32   0", float32(0.0), reflect.Float32)
	test("float32 +pi", float32(3.141529), reflect.Float32)

	test("float64 -pi", -3.141529, reflect.Float64)
	test("float64   0", 0.0, reflect.Float64)
	test("float64 +pi", 3.141529, reflect.Float64)

	test("int-1", -1, reflect.Int)
	test("int0", 0, reflect.Int)
	test("int+1", +1, reflect.Int)

	test("empty string", "", reflect.String)
	test("test string", "test string", reflect.String)

}
