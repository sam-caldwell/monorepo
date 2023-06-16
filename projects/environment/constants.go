package environment

const (
	intVarSize   = 64
	intBase      = 10
	floatVarSize = 64

	defaultBoolValue   = false
	defaultFloatValue  = 0.0
	defaultIntValue    = 0
	defaultStringValue = ""

	errEnvVarNotFound     = "required value not found"
	errEnvBoundCheck      = "value out of bounds"
	errReadingValue       = "error parsing value (%v)"
	errPatternCheckFailed = "pattern check failed"
)
