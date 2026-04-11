package keyvalue

import (
	"fmt"
	"testing"
)

func TestKeyValue_Exists(t *testing.T) {
	var kv KeyValue[string, int]
	kv.data = make(map[string]int)
	kv.lock.Lock()
	kv.lock.Unlock()
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("test%d", i)
		value := i
		kv.data[key] = value
	}
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("test%d", i)
		if !kv.Exists(key) {
			t.Fatalf("key %s not exists", key)
		}
	}

}
