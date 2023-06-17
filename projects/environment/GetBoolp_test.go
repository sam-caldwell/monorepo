package environment

import (
	"fmt"
	"testing"
)

func TestGetBoolp(t *testing.T) {
	runBoolpTest(t, "b0", GetBoolp, "true", nil, true)
	runBoolpTest(t, "b1", GetBoolp, "false", nil, false)
	runBoolpTest(t, "b2", GetBoolp, "TRUE", nil, true)
	runBoolpTest(t, "b3", GetBoolp, "FALSE", nil, false)
	runBoolpTest(t, "b4", GetBoolp, "T", nil, true)
	runBoolpTest(t, "b5", GetBoolp, "F", nil, false)
	runBoolpTest(t, "b6", GetBoolp, "t", nil, true)
	runBoolpTest(t, "b7", GetBoolp, "f", nil, false)
	runBoolpTest(t, "b8", GetBoolp, "1", nil, true)
	runBoolpTest(t, "b9", GetBoolp, "0", nil, false)

	errBadInput := fmt.Errorf("strconv.ParseBool: parsing \"BadString\": invalid syntax")
	runBoolpTest(t, "b10", GetBoolp, "BadString", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"1.0\": invalid syntax")
	runBoolpTest(t, "b11", GetBoolp, "1.0", errBadInput, false)

	errBadInput = fmt.Errorf("strconv.ParseBool: parsing \"0.0\": invalid syntax")
	runBoolpTest(t, "b12", GetBoolp, "0.0", errBadInput, false)
}
