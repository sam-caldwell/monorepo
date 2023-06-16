package environment

import (
	"fmt"
	"testing"
)

func TestRequireBoolp(t *testing.T) {
	runBoolpTest(t, "b0", RequireBoolp, "true", nil, true)
	runBoolpTest(t, "b1", RequireBoolp, "false", nil, false)
	runBoolpTest(t, "b2", RequireBoolp, "TRUE", nil, true)
	runBoolpTest(t, "b3", RequireBoolp, "FALSE", nil, false)
	runBoolpTest(t, "b4", RequireBoolp, "T", nil, true)
	runBoolpTest(t, "b5", RequireBoolp, "F", nil, false)
	runBoolpTest(t, "b6", RequireBoolp, "t", nil, true)
	runBoolpTest(t, "b7", RequireBoolp, "f", nil, false)
	runBoolpTest(t, "b8", RequireBoolp, "1", nil, true)
	runBoolpTest(t, "b9", RequireBoolp, "0", nil, false)

	errBadInput := fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	runBoolpTest(t, "b10", RequireBoolp, "BadString", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	runBoolpTest(t, "b11", RequireBoolp, "1.0", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	runBoolpTest(t, "b12", RequireBoolp, "0.0", errBadInput, false)
}
