package environment_testing

import "testing"

func RunRequiredFloatpTest(t *testing.T, name string, f TestFloatFunc, err error) {
	const envVar = "TestRequiredFloatp"
	const emptyString = ""

	Setup(t, envVar, emptyString) //Make sure it does exist (as an invalid empty string)
	if _, actualError := f(name); actualError == nil {
		checkError(t, &name, actualError, err)
	}
	TearDown(t, envVar) //Make sure it doesn't exist
	if _, actualError := f(name); actualError == nil {
		checkError(t, &name, actualError, err)
	}
}
