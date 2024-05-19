package keyvalue

// OverwriteProtection - Used by KeyValue.Initialize() to reallocate key-value map, if already allocated
//
//	(c) 2023 Sam Caldwell.  MIT License
type OverwriteProtection bool

const (
	// overwrite - if KeyValue is already allocated, reallocate the memory and call gc
	overwrite OverwriteProtection = true
	// preserve - if KeyValue is already allocated, preserve the existing state.
	preserve OverwriteProtection = false
)

// Column parsing values
const (
	//columnCount - number of columns we will parse data for to convert to KeyValue
	columnCount = 2
	//keyColumn - the index of the parsed key value
	keyColumn = 0
	//valueColumn - the starting index of the parsed value (1...n will be used)
	valueColumn = 1
)
