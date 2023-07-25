package convert

import "testing"

func TestMorseCodeMap(t *testing.T) {
	for morse, ascii := range morseCodeToASCII {
		actualMorse, ok := AsciiToMorseCode[ascii]
		if !ok {
			t.Errorf("Character %q is not present in the morseCode map", ascii)
			continue
		}

		if actualMorse != morse {
			t.Errorf("Invalid Morse code mapping for character %q. Expected: %q, Got: %q",
				ascii, morse, actualMorse)
		}
	}
}
