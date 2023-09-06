package vartype

import "testing"

func TestGeneric_Set(t *testing.T) {
	var o Generic

	test := func(name string, data any) {
		o.Set(data)
		if o.data != data {
			t.Fatalf("Failed on %s", name)
		}
	}

	test("bool true", true)
	test("bool false", false)

	test("float32 -3.141529", float32(-3.141529))
	test("float32 0", float32(0))
	test("float32 3.141529", float32(3.141529))

	test("float64 -3.141529", -3.141529)
	test("float64 0", float64(0))
	test("float64 3.141529", 3.141529)

	test("int -1", -1)
	test("int  0", 0)
	test("int +1", +1)

	test("string", "test")
	test("string", "")

	test("nil value", nil)

}
