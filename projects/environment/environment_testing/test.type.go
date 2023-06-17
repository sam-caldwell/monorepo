package environment_testing

// TestBoolFunc - function pointer for Bool
type TestBoolFunc func(arg string) (bool, error)

// TestBoolpFunc - function pointer for Boolp
type TestBoolpFunc func(arg *string) (bool, error)

// TestFloatFunc - function pointer for Float
type TestFloatFunc func(arg string) (float64, error)

// TestFloatpFunc - function pointer for Float
type TestFloatpFunc func(arg *string) (float64, error)

// TestFloatSanityFunc - function pointer for Float
type TestFloatSanityFunc func(arg string, min float64, max float64) (float64, error)

// TestFloatpSanityFunc - function pointer for Float
type TestFloatpSanityFunc func(arg *string, min float64, max float64) (float64, error)

// TestIntFunc - function pointer for Int
type TestIntFunc func(arg string) (int, error)

// TestIntpFunc - function pointer for Int
type TestIntpFunc func(arg *string) (int, error)

// TestIntSanityFunc - function pointer for int
type TestIntSanityFunc func(arg string, min int, max int) (int, error)

// TestIntpSanityFunc - function pointer for int
type TestIntpSanityFunc func(arg *string, min int, max int) (int, error)

// TestStringFunc - function pointer for strings
type TestStringFunc func(arg string) (string, error)

// TestStringpFunc - function pointer for stringp
type TestStringpFunc func(arg *string) (string, error)

// TestStringSanityFunc - function pointer for string
type TestStringSanityFunc func(arg string, pattern string) (string, error)

// TestStringpSanityFunc - function pointer for string
type TestStringpSanityFunc func(arg *string, pattern *string) (string, error)
