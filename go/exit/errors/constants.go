package errors

const (
	// Details - Modifier (any) to add Details to an error
	Details = "(%v)"

	// DuplicateEntry - indicates that an operation failed because a duplicate was created.
	DuplicateEntry             = "duplicate entry"
	IndexOutOfRange            = "index out of range"
	InternalError              = "internal error"
	InvalidCommand             = "unknown command"
	InvalidCommandWithDetail   = InvalidCommand + Details
	InvalidContextId           = "invalid contextId"
	InvalidContextIdWithDetail = InvalidContextId + Details
	InvalidInput               = "invalid input"
	LockCheckFailed            = "lock check failed"
	MalformedGuid              = "malformed GUID"
	//MissingArguments - One or more expected inputs are missing
	MissingArguments           = "missing argument"
	MissingArgWithDetail       = MissingArguments + Details
	MissingColor               = "Missing color"
	MissingContextId           = "missing contextId (uuid)"
	MissingField               = "missing field"
	NotFound                   = "not found"
	NotInitialized             = "not initialized"
	OverflowError              = "overflow error"
	ReadOnly                   = "read only"
	TypeMismatch               = "type mismatch"
	UnableToDetectFamily       = "unable to detect operating system family"
	UnknownCommand             = "Unknown command"
	UnsupportedOpsys           = "unsupported operating system"
	UnsupportedCpuArchitecture = "unsupported cpu architecture"
	UnsupportedLanguage        = "unsupported language"
	UnsupportedVersion         = "unsupported version"
)
