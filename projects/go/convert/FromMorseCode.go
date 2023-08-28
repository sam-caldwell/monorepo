package convert

/*
 * convert.FromMorseCode()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * A simple function that converts a morse code string
 * into an ASCII string.
 */
import (
	"fmt"
	ansi "github.com/sam-caldwell/monorepo/v2/projects/go/ansi"
	"strings"
)

func FromMorseCode(morseCode string) (string, error) {
	morseCode = strings.TrimSpace(morseCode)
	morseWords := strings.Split(morseCode, "/")
	output := ""

	for wordPos, word := range morseWords {
		morseChars := strings.Split(word, " ")
		for charPos, morseChar := range morseChars {
			if asciiChar, ok := morseCodeToASCII[morseChar]; ok {
				ansi.Green().Printf("%d:%d[%s] ", wordPos, charPos, asciiChar)
				output += asciiChar
			} else {
				if strings.TrimSpace(morseChar) != "" {
					return output, fmt.Errorf("invalid morseChar: '%s'", morseChar)
				}
			}
		}
		output += " "
	}
	ansi.LF().Reset()

	return strings.TrimSpace(output), nil
}
