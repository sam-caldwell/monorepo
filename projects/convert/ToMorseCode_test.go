package convert

import (
	"testing"
)

func TestToMorseCode(t *testing.T) {
	input := "Hello, World!"
	expectedOutput := ".... . .-.. .-.. --- --..-- / .-- --- .-. .-.. -.. -.-.--"

	output, err := ToMorseCode(input)
	if err != nil {
		t.Fatal(err)
	}

	if output != expectedOutput {
		t.Errorf("ToMorseCode() test failed: expected '%s', got '%s'", expectedOutput, output)
	}

	t.Run("EmptyString", func(t *testing.T) {
		input := ""
		expectedOutput := ""

		output, err := ToMorseCode(input)
		if err != nil {
			t.Fatal(err)
		}

		if output != expectedOutput {
			t.Errorf("ToMorseCode() test failed: expected '%s', got '%s'", expectedOutput, output)
		}
	})

	t.Run("OnlySpaces", func(t *testing.T) {
		input := "    "
		expectedOutput := "/ / / /"

		output, err := ToMorseCode(input)
		if err != nil {
			t.Fatal(err)
		}

		if output != expectedOutput {
			t.Errorf("ToMorseCode() test failed: expected '%s', got '%s'", expectedOutput, output)
		}
	})
}
