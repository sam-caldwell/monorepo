package environment_testing

import "testing"

func RunRequiredIntTest(t *testing.T, name string, f TestFloatFunc, err error) {
	const envVar = "TestRequiredInt"
	const emptyString = ""

	Setup(t, envVar, emptyString) //Make sure it does exist (as an invalid empty string)
	if _, actualError := f(name); actualError == nil {
		errorCheck(t, &name, actualError, err)
	}
	TearDown(t, envVar) //Make sure it doesn't exist
	if _, actualError := f(name); actualError == nil {
		errorCheck(t, &name, actualError, err)
	}
}
