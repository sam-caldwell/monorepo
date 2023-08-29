package hijack

//func TestRequireCompatibleFunctions(t *testing.T) {
//	// Define sample functions
//	functionA := func() {
//		// Function body
//	}
//
//	functionB := func() {
//		// Function body
//	}
//
//	// Call the function under test
//	err := requireCompatibleFunctions(functionA, functionB)
//
//	// Check if an error occurred
//	if err != nil {
//		t.Errorf("Expected compatible functions, but got an error: %v", err)
//	}
//
//	// Define a non-function object
//	nonFunction := "This is not a function"
//
//	// Call the function under test with a non-function object
//	err = requireCompatibleFunctions(functionA, nonFunction)
//
//	// Check if an error occurred
//	expectedErrorMessage := fmt.Sprintf("%s has to be a function", getNameOfFunction(nonFunction))
//	if err == nil || err.Error() != expectedErrorMessage {
//		t.Errorf("Expected error: %s, but got: %v", expectedErrorMessage, err)
//	}
//
//	// Define a different function type
//	functionC := func(i int) {
//		// Function body
//	}
//
//	// Call the function under test with incompatible function types
//	err = requireCompatibleFunctions(functionA, functionC)
//
//	// Check if an error occurred
//	expectedErrorMessage = fmt.Sprintf("both %s and %s must be of the same type", getNameOfFunction(functionA), getNameOfFunction(functionC))
//	if err == nil || err.Error() != expectedErrorMessage {
//		t.Errorf("Expected error: %s, but got: %v", expectedErrorMessage, err)
//	}
//}
