package keyvalue

/*
 * keyvalue.Initialize()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Allocate memory for our Key-value map
 */

// Initialize - Allocate memory for our Key-value map
func (kv *KeyValue) Initialize(sz int, allowOverwrite OverwriteProtection) {
	//
	// Only allocate memory if not allocated or 'overwrite' flag is set.
	//
	if (kv.data == nil) || allowOverwrite {
		//A less-than-zero length memory allocation sounds like a fantasy from Dilber's Mr. PointyHead
		//Let's not allow that.
		if sz < 0 {
			sz = 0
		}
		kv.data = make(Map, sz)
	}
}
