package environment_testing

import "testing"

func RunRequiredIntpTest(t *testing.T, name string, f TestIntpFunc, err error) {
	const envVar = "TestRequiredIntp"
	const emptyString = ""

	Setup(t, envVar, emptyString) //Make sure it does exist (as an invalid empty string)
	if _, actualError := f(&name); actualError == nil {
		checkError(t, &name, actualError, err)
	}
	TearDown(t, envVar) //Make sure it doesn't exist
	if _, actualError := f(&name); actualError == nil {
		checkError(t, &name, actualError, err)
	}
}
