package errors

const (
	// Details - Modifier (any) to add Details to an error
	Details = "(%v)"

	// DuplicateEntry - indicates that an operation failed because a duplicate was created.
	DuplicateEntry = "duplicate entry"

	//MissingArguments - One or more expected inputs are missing
	MissingArguments        = "missing argument"
	ErrMissingArgWithDetail = MissingArguments + Details

	ErrIndexOutOfRange            = "index out of range"
	ErrInternalError              = "internal error"
	ErrInvalidCommand             = "unknown command"
	ErrInvalidCommandWithDetail   = ErrInvalidCommand + Details
	ErrInvalidContextId           = "Invalid contextId"
	ErrInvalidContextIdWithDetail = ErrInvalidContextId + Details
	ErrInvalidInput               = "invalid input"
	ErrLockCheckFailed            = "lock check failed"
	ErrMissingColor               = "Missing color"
	ErrMissingContextId           = "missing contextId (uuid)"
	ErrNotFound                   = "not found"
	ErrNotInitialized             = "not initialized"
	ErrTypeMismatch               = "type mismatch"
	ErrUnknownCommand             = "Unknown command"

	ErrUnsupportedOpsys             = "unsupported operating system"
	ErrUnsupportedOpsysWithDetail   = ErrUnsupportedOpsys + Details
	ErrUnsupportedVersion           = "unsupported version"
	ErrUnsupportedVersionWithDetail = ErrUnsupportedVersion + Details
)
