package errors

const (
	ArraySizeError   = "array size error"
	BoundsCheckError = "bounds check error"
	CannotOpenFile   = "cannot open file"
	CannotReadFile   = "cannot read file"
	// Details - Modifier (any) to add Details to an error
	Details                    = "(%v)"
	DuplicateEntry             = "duplicate entry"
	EmptySet                   = "empty set"
	IndexOutOfRange            = "index out of range"
	InternalError              = "internal error"
	InvalidCommand             = "unknown command"
	InvalidContextId           = "invalid contextId"
	InvalidInput               = "invalid input"
	LockCheckFailed            = "lock check failed"
	MalformedGuid              = "malformed GUID"
	MissingArguments           = "missing argument"
	MissingColor               = "Missing color"
	MissingContextId           = "missing contextId (uuid)"
	MissingField               = "missing field"
	NotFound                   = "not found"
	NotInitialized             = "not initialized"
	OverflowError              = "overflow error"
	ReadOnly                   = "read only"
	TypeMismatch               = "type mismatch"
	UnableToDetectFamily       = "unable to detect operating system family"
	UninitializedValue         = "uninitialized value"
	UnknownCommand             = "Unknown command"
	UnsupportedType            = "Unsupported type"
	UnsupportedOpsys           = "unsupported operating system"
	UnsupportedCpuArchitecture = "unsupported cpu architecture"
	UnsupportedLanguage        = "unsupported language"
	UnsupportedVersion         = "unsupported version"
	ValueTooSmall              = "value too small"
)
