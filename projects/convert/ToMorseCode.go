package convert

import (
	"fmt"
	"strings"
)

// ToMorseCode - Write an ASCII string in morse code.
//
//	(c) 2023 Sam Caldwell.  MIT License
func ToMorseCode(input string) (string, error) {
	output := ""

	for _, char := range input {
		upperChar := strings.ToUpper(string(char))
		if morseCodeValue, ok := AsciiToMorseCode[upperChar]; ok {
			output += morseCodeValue + " "
		} else if upperChar == " " {
			output += " / "
		} else {
			return output, fmt.Errorf("invalid character (%d)", char)
		}
	}
	return strings.TrimSpace(output), nil
}
