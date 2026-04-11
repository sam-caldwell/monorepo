package convert

import (
	"testing"
)

func TestFromMorseCode(t *testing.T) {
	morseCode := ".... . .-.. .-.. --- --..-- / .-- --- .-. .-.. -.. -.-.--"
	expectedOutput := "HELLO, WORLD!"

	output, err := FromMorseCode(morseCode)

	if err != nil {
		t.Fatalf("FromMorseCode() test failed with error: %v", err)
	}

	if output != expectedOutput {
		t.Fatalf("FromMorseCode() test failed: expected '%s', got '%s'", expectedOutput, output)
	}

	t.Run("EmptyString", func(t *testing.T) {
		morseCode := ""
		expectedOutput := ""

		output, err := FromMorseCode(morseCode)

		if err != nil {
			t.Fatalf("FromMorseCode() test failed with error: %v", err)
		}

		if output != expectedOutput {
			t.Fatalf("FromMorseCode() test failed: expected '%s', got '%s'", expectedOutput, output)
		}
	})

	t.Run("InvalidMorseCode", func(t *testing.T) {
		morseCode := "This is Invalid"
		expectedError := "invalid morseChar: 'This'"

		output, err := FromMorseCode(morseCode)

		if err == nil {
			t.Error("FromMorseCode() test failed: expected error, got nil")
		}

		if err.Error() != expectedError {
			t.Fatalf("FromMorseCode() test failed: expected error '%s', got '%s'", expectedError, err.Error())
		}

		if output != "" {
			t.Fatalf("FromMorseCode() test failed: expected empty output, got '%s'", output)
		}
	})
}
