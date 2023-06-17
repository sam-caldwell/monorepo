package environment

import (
	"fmt"
	"testing"
)

func TestSanitizeBool(t *testing.T) {
	runBoolTest(t, "b0", SanitizeBool, "true", nil, true)
	runBoolTest(t, "b1", SanitizeBool, "false", nil, false)
	runBoolTest(t, "b2", SanitizeBool, "TRUE", nil, true)
	runBoolTest(t, "b3", SanitizeBool, "FALSE", nil, false)
	runBoolTest(t, "b4", SanitizeBool, "T", nil, true)
	runBoolTest(t, "b5", SanitizeBool, "F", nil, false)
	runBoolTest(t, "b6", SanitizeBool, "t", nil, true)
	runBoolTest(t, "b7", SanitizeBool, "f", nil, false)
	runBoolTest(t, "b8", SanitizeBool, "1", nil, true)
	runBoolTest(t, "b9", SanitizeBool, "0", nil, false)

	errBadInput := fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	runBoolTest(t, "b10", SanitizeBool, "BadString", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	runBoolTest(t, "b11", SanitizeBool, "1.0", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	runBoolTest(t, "b12", SanitizeBool, "0.0", errBadInput, false)
}
