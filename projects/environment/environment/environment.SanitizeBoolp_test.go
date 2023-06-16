package environment

import (
	"fmt"
	"testing"
)

func TestSanitizeBoolp(t *testing.T) {
	runBoolpTest(t, "b0", SanitizeBoolp, "true", nil, true)
	runBoolpTest(t, "b1", SanitizeBoolp, "false", nil, false)
	runBoolpTest(t, "b2", SanitizeBoolp, "TRUE", nil, true)
	runBoolpTest(t, "b3", SanitizeBoolp, "FALSE", nil, false)
	runBoolpTest(t, "b4", SanitizeBoolp, "T", nil, true)
	runBoolpTest(t, "b5", SanitizeBoolp, "F", nil, false)
	runBoolpTest(t, "b6", SanitizeBoolp, "t", nil, true)
	runBoolpTest(t, "b7", SanitizeBoolp, "f", nil, false)
	runBoolpTest(t, "b8", SanitizeBoolp, "1", nil, true)
	runBoolpTest(t, "b9", SanitizeBoolp, "0", nil, false)

	errBadInput := fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	runBoolpTest(t, "b10", SanitizeBoolp, "BadString", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	runBoolpTest(t, "b11", SanitizeBoolp, "1.0", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	runBoolpTest(t, "b12", SanitizeBoolp, "0.0", errBadInput, false)
}
