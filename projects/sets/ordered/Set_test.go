package ordered

import "testing"

func TestSet_Struct(t *testing.T) {
	tests := []struct {
		name string
		typ  interface{}
	}{
		{"int", int(0)},
		{"int8", int8(0)},
		{"int16", int16(0)},
		{"int32", int32(0)},
		{"int64", int64(0)},
		{"uint", uint(0)},
		{"uint8", uint8(0)},
		{"uint16", uint16(0)},
		{"uint32", uint32(0)},
		{"uint64", uint64(0)},
		{"rune", rune(0)},
		{"string", ""},
		{"bool", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s Set[interface{}]
			if s.data != nil {
				t.Fail()
			}
		})
	}
}
