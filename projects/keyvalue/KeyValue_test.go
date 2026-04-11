package keyvalue

import (
	"testing"
)

func TestKeyValue_struct(t *testing.T) {
	t.Run("initial state", func(t *testing.T) {
		var kv KeyValue[int, int]
		if kv.data != nil {
			t.Fatal("initial state should be nil")
		}
	})
	t.Run("verify keyValue can be different types", func(t *testing.T) {
		var kv1 KeyValue[string, string]
		kv1.Initialize(0, false)
		var kv2 KeyValue[string, []byte]
		kv2.Initialize(0, false)
		var kv3 KeyValue[string, bool]
		kv3.Initialize(0, false)
		var kv4 KeyValue[string, int]
		kv4.Initialize(0, false)
		var kv5 KeyValue[string, int8]
		kv5.Initialize(0, false)
		var kv6 KeyValue[string, int16]
		kv6.Initialize(0, false)
		var kv7 KeyValue[string, int32]
		kv7.Initialize(0, false)
		var kv8 KeyValue[string, int64]
		kv8.Initialize(0, false)
		var kv9 KeyValue[string, uint]
		kv9.Initialize(0, false)
		var kvA KeyValue[string, uint8]
		kvA.Initialize(0, false)
		var kvB KeyValue[string, uint16]
		kvB.Initialize(0, false)
		var kvC KeyValue[string, uint32]
		kvC.Initialize(0, false)
		var kvD KeyValue[string, uint64]
		kvD.Initialize(0, false)
		var kvE KeyValue[string, []byte]
		kvE.Initialize(0, false)
		var kvF KeyValue[int, bool]
		kvF.Initialize(0, false)
		var kvG KeyValue[int, int]
		kvG.Initialize(0, false)
		var kvH KeyValue[int, int8]
		kvH.Initialize(0, false)
		var kvI KeyValue[int, int16]
		kvI.Initialize(0, false)
		var kvJ KeyValue[int, int32]
		kvJ.Initialize(0, false)
		var kvK KeyValue[int, int64]
		kvK.Initialize(0, false)
		var kvL KeyValue[int, uint]
		kvL.Initialize(0, false)
		var kvM KeyValue[int, uint8]
		kvM.Initialize(0, false)
		var kvN KeyValue[int, uint16]
		kvN.Initialize(0, false)
		var kvO KeyValue[int, uint32]
		kvO.Initialize(0, false)
		var kvP KeyValue[int, uint64]
		kvP.Initialize(0, false)
	})
}
