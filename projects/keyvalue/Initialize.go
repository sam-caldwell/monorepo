package keyvalue

import "runtime"

// Initialize - Allocate memory for our Key-value map
//
//	If allowOverwrite is passed in, then the memory for the underlying map will be reallocated,
//	and the method will force garbage collection to run.  This is intended to ensure information
//	is not left in memory longer than needed for security reasons, accepting the performance hit.
//
//	If sz (size) is less than zero, the value is assumed to be 0.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) Initialize(sz int, allowOverwrite OverwriteProtection) {
	//
	// Only allocate memory if not allocated or 'overwrite' flag is set.
	//
	if allowOverwrite {
		kv.lock.Lock()
		kv.data = nil
		//Force garbage collection in case of reallocation to ensure we don't leave stuff in memory.
		runtime.GC()
		kv.lock.Unlock()
	}
	if kv.data == nil {
		//A less-than-zero length memory allocation sounds like a fantasy from Dilber's Mr. PointyHead
		//Let's not allow that.
		if sz < 0 {
			sz = 0
		}
		kv.lock.Lock()
		kv.data = make(map[KeyType]ValueType, sz)
		kv.lock.Unlock()
	}
}
