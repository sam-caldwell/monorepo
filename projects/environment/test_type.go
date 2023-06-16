package environment

// testBoolFunc - function pointer for Bool
type testBoolFunc func(arg string) (bool, error)

// testBoolpFunc - function pointer for Boolp
type testBoolpFunc func(arg *string) (bool, error)

// testFloatFunc - function pointer for Float
type testFloatFunc func(arg string) (float64, error)

// testFloatpFunc - function pointer for Float
type testFloatpFunc func(arg *string) (float64, error)

// testFloatSanityFunc - function pointer for Float
type testFloatSanityFunc func(arg string, min float64, max float64) (float64, error)

// testFloatpSanityFunc - function pointer for Float
type testFloatpSanityFunc func(arg *string, min float64, max float64) (float64, error)

// testIntFunc - function pointer for Int
type testIntFunc func(arg string) (int, error)

// testIntpFunc - function pointer for Int
type testIntpFunc func(arg *string) (int, error)

// testIntSanityFunc - function pointer for int
type testIntSanityFunc func(arg string, min int, max int) (int, error)

// testIntpSanityFunc - function pointer for int
type testIntpSanityFunc func(arg *string, min int, max int) (int, error)

// testStringFunc - function pointer for strings
type testStringFunc func(arg string) (string, error)

// testStringpFunc - function pointer for stringp
type testStringpFunc func(arg *string) (string, error)

// testStringSanityFunc - function pointer for string
type testStringSanityFunc func(arg string, pattern string) (string, error)

// testStringpSanityFunc - function pointer for string
type testStringpSanityFunc func(arg *string, pattern *string) (string, error)
