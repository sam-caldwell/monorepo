package keyvalue

/*
 * KeyValue (constants)
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 */

type OverwriteProtection bool

// Used by KeyValue.Initialize() to reallocate key-value map, if already allocated
const (
	overwrite OverwriteProtection = true
	preserve  OverwriteProtection = false
)

// Column parsing values
const (
	columnCount = 2
	keyColumn   = 0
	valueColumn = 1
)
