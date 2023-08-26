package convert

import (
	"testing"
)

func TestFromMorseCode(t *testing.T) {
	morseCode := ".... . .-.. .-.. --- --..-- / .-- --- .-. .-.. -.. -.-.--"
	expectedOutput := "HELLO, WORLD!"

	output, err := FromMorseCode(morseCode)

	if err != nil {
		t.Errorf("FromMorseCode() test failed with error: %v", err)
	}

	if output != expectedOutput {
		t.Errorf("FromMorseCode() test failed: expected '%s', got '%s'", expectedOutput, output)
	}

	t.Run("EmptyString", func(t *testing.T) {
		morseCode := ""
		expectedOutput := ""

		output, err := FromMorseCode(morseCode)

		if err != nil {
			t.Errorf("FromMorseCode() test failed with error: %v", err)
		}

		if output != expectedOutput {
			t.Errorf("FromMorseCode() test failed: expected '%s', got '%s'", expectedOutput, output)
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
			t.Errorf("FromMorseCode() test failed: expected error '%s', got '%s'", expectedError, err.Error())
		}

		if output != "" {
			t.Errorf("FromMorseCode() test failed: expected empty output, got '%s'", output)
		}
	})
}
