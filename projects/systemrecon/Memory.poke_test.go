//go:build darwin || linux || windows
// +build darwin linux windows

package systemrecon

//func TestPoke(t *testing.T) {
//var err error
//func() {
//	/*
//	 * Set up the test
//	 */
//	const (
//		initialValue = "initial state"
//		alteredValue = "altered state"
//	)
//
//	var actualValue = words.EmptyString
//	// Test case: Write data to memory with no error
//	targetFunc := func() { actualValue = initialValue }
//	alteredFunc := func() { actualValue = alteredValue }
//
//	targetFuncPtr := uintptr(unsafe.Pointer(&targetFunc))
//	alteredFuncPtr := uintptr(unsafe.Pointer(&alteredFunc))
//
//	// Execute the original targetFunc()
//	targetFunc()
//	if actualValue != initialValue {
//		t.Fatal("expected initialValue")
//	}
//	// Save the altered state
//	alteredBytes := peek(alteredFuncPtr, 16)
//	/*
//	 * overwrite targetFuncPtr
//	 */
//	//Overwrite targetFunc with alteredFunc
//	if err = poke(targetFuncPtr, alteredBytes); err != nil {
//		t.Fatalf("Unexpected error: %v", err)
//	}
//	targetFunc()
//	if actualValue != alteredValue {
//		t.Fatal("expected alteredValue")
//	}
//}()
//
//func() {
//	// Test case: Write empty data to memory with no error
//	emptyMemoryAddress := uintptr(0x2000)
//	var emptySourceData []byte
//	err = poke(emptyMemoryAddress, emptySourceData)
//
//	// Verify that no error occurred
//	if err != nil {
//		t.Fatalf("Unexpected error: %v", err)
//	}
//}()
//}
