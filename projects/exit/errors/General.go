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
	InvalidContextId           = "Invalid contextId"
	InvalidContextIdWithDetail = InvalidContextId + Details
	InvalidInput               = "invalid input"
	LockCheckFailed            = "lock check failed"
	//MissingArguments - One or more expected inputs are missing
	MissingArguments             = "missing argument"
	MissingArgWithDetail         = MissingArguments + Details
	MissingColor                 = "Missing color"
	MissingContextId             = "missing contextId (uuid)"
	NotFound                     = "not found"
	NotInitialized               = "not initialized"
	TypeMismatch                 = "type mismatch"
	UnknownCommand               = "Unknown command"
	UnsupportedOpsys             = "unsupported operating system"
	UnsupportedOpsysWithDetail   = UnsupportedOpsys + Details
	UnsupportedVersion           = "unsupported version"
	UnsupportedVersionWithDetail = UnsupportedVersion + Details
)
