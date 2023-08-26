package systemrecon

//func TestPeek(t *testing.T) {
//	const (
//		NULL          = 0x00
//		expectedValue = "This is a test string of several words"
//	)
//
//	func() {
//		// Test case: Peek into a memory address with valid data
//
//		data := []byte(expectedValue)
//		addr := uintptr(unsafe.Pointer(&data[0]))
//		length := len(data)
//		actualValue := Peek(addr, length)
//		// Verify that the peeked data matches the original data
//		if string(*actualValue) != expectedValue {
//			t.Errorf("Peeked data does not match the original data")
//		}
//	}()
//
//	func() {
//		// Test case: Peek into a memory address with zero length
//		var emptyData []byte
//		emptyAddr := uintptr(unsafe.Pointer(&emptyData))
//		emptyLength := len(emptyData)
//		actualValue := Peek(emptyAddr, emptyLength)
//
//		// Verify that peeking into an empty memory address returns an empty slice
//		if len(*actualValue) != emptyLength {
//			t.Errorf("Peeked data is not empty for an empty memory address")
//		}
//	}()
//
//	func() {
//		// Test case: send in a null pointer and zero length
//		emptyAddr := uintptr(NULL)
//		emptyLength := 10
//		actualValue := Peek(emptyAddr, emptyLength)
//
//		// Verify that peeking into an empty memory address returns an empty slice
//		if len(*actualValue) != emptyLength {
//			t.Errorf("Peeked data is not empty for an empty memory address\n"+
//				"actualValue size:  %d\n"+
//				"     actualValue: '%0x'", len(*actualValue), actualValue)
//		}
//	}()
//}
