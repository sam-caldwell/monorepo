package keyvalue

import "testing"

func TestKeyValue_Set(t *testing.T) {
	var kv KeyValue[string, string]
	t.Run("Set uninitialized state", func(t *testing.T) {
		kv.Set("foo", "bar")
		if value, ok := kv.data["foo"]; !ok || value != "bar" {
			t.Fatalf("Set uninitialized key")
		}
	})
	t.Run("Set initialized state", func(t *testing.T) {
		kv.Set("moo", "mar")
		if value, ok := kv.data["foo"]; !ok || value != "bar" {
			t.Fatalf("Set uninitialized key (foo)")
		}
		if value, ok := kv.data["moo"]; !ok || value != "mar" {
			t.Fatalf("Set uninitialized key (moo)")
		}
	})
}

//
//func TestSetBoolHappyPath(t *testing.T) {
//	kv := KeyValue{}
//
//	kv.SetBool("testkey", true)
//
//	if val, ok := kv.data["testkey"].(bool); !ok {
//		t.Fatalf("Key was not found in map or is not of the correct type")
//	} else if val != true {
//		t.Fatalf("Incorrect value. Expected true, but got %v", val)
//	}
//}
//
//func TestSetBoolOverwriteExistingKey(t *testing.T) {
//	kv := KeyValue{
//		data: Map{
//			"testkey": false,
//		},
//	}
//
//	kv.SetBool("testkey", true)
//
//	if val, ok := kv.data["testkey"].(bool); !ok {
//		t.Fatalf("Key was not found in map or is not of the correct type")
//	} else if val != true {
//		t.Fatalf("Incorrect value. Expected true, but got %v", val)
//	}
//}
//
//func TestSetBoolWithNilMap(t *testing.T) {
//	kv := KeyValue{}
//
//	kv.SetBool("testkey", true)
//
//	if val, ok := kv.data["testkey"].(bool); !ok {
//		t.Fatalf("Key was not found in map or is not of the correct type")
//	} else if val != true {
//		t.Fatalf("Incorrect value. Expected true, but got %v", val)
//	}
//}
