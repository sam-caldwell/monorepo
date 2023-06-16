package environment

import (
	"fmt"
	"testing"
)

func TestGetBool(t *testing.T) {
	runBoolTest(t, "b0", GetBool, "true", nil, true)
	runBoolTest(t, "b1", GetBool, "false", nil, false)
	runBoolTest(t, "b2", GetBool, "TRUE", nil, true)
	runBoolTest(t, "b3", GetBool, "FALSE", nil, false)
	runBoolTest(t, "b4", GetBool, "T", nil, true)
	runBoolTest(t, "b5", GetBool, "F", nil, false)
	runBoolTest(t, "b6", GetBool, "t", nil, true)
	runBoolTest(t, "b7", GetBool, "f", nil, false)
	runBoolTest(t, "b8", GetBool, "1", nil, true)
	runBoolTest(t, "b9", GetBool, "0", nil, false)

	errBadInput := fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	runBoolTest(t, "b10", GetBool, "BadString", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	runBoolTest(t, "b11", GetBool, "1.0", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	runBoolTest(t, "b12", GetBool, "0.0", errBadInput, false)
}
