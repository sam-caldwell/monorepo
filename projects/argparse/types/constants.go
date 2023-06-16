package types

const (

	// Boolean - Represents a boolean (true/false)
	Boolean ArgTypes = 2

	// Flag - Represents a boolean flag (present - true, absent - false)
	Flag ArgTypes = 4

	// Float - Represents a 64-bit floating-point number
	Float ArgTypes = 3

	// Integer - Represents a 64-bit integer
	Integer ArgTypes = 1

	// String - Represents a string
	String ArgTypes = 0
)

const (
	// eMsgTypeCheck - type check error string
	eMsgTypeCheck = "type-check failed (%s)"

	// eMsgTypeCheckUnknown - type-check failed error
	eMsgTypeCheckUnknown = "type-check failed (unknown type)"

	errInvalidArgType = "invalid descriptor type %s"

	errUnknownType = "unknown-type(%d)"
)
