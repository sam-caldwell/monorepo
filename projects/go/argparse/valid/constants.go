package valid

const (
	validArgRegex      = `^([a-zA-Z]|[a-zA-Z][a-zA-Z0-9_-]*[a-zA-Z0-9])$`
	validShortArgRegex = `^-[a-z0-9]$`
	validLongArgRegex  = `^--[a-zA-Z][a-zA-Z0-9_-]*[a-zA-Z0-9]$`
	validYearRegex     = `^[0-9]{4}$`
)

const (
	errEmptyOrNilObject = "missing or nil input"
	errExpectedShortArg = "expected short descriptor for '%s'"
	errExpectedLongArg  = "expected long descriptor for '%s'"
	errExpectedNameArg  = "expected name descriptor for '%s'"

	errInvalidYear = "invalid year %d"
)
