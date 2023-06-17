package environment

import (
	"fmt"
	"testing"
)

func TestRequireBool(t *testing.T) {
	runBoolTest(t, "b0", RequireBool, "true", nil, true)
	runBoolTest(t, "b1", RequireBool, "false", nil, false)
	runBoolTest(t, "b2", RequireBool, "TRUE", nil, true)
	runBoolTest(t, "b3", RequireBool, "FALSE", nil, false)
	runBoolTest(t, "b4", RequireBool, "T", nil, true)
	runBoolTest(t, "b5", RequireBool, "F", nil, false)
	runBoolTest(t, "b6", RequireBool, "t", nil, true)
	runBoolTest(t, "b7", RequireBool, "f", nil, false)
	runBoolTest(t, "b8", RequireBool, "1", nil, true)
	runBoolTest(t, "b9", RequireBool, "0", nil, false)

	errBadInput := fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	runBoolTest(t, "b10", RequireBool, "BadString", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	runBoolTest(t, "b11", RequireBool, "1.0", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	runBoolTest(t, "b12", RequireBool, "0.0", errBadInput, false)
}
