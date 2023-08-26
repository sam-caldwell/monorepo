package convert

/*
 * convert.ToMorseCode()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * A simple function that converts a message string to
 * morse code and returns the result as dots and dashes.
 */
import (
	"fmt"
	"strings"
)

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
