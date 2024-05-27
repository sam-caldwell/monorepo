package convert

import (
	"fmt"
	ansi2 "github.com/sam-caldwell/monorepo/go/ansi"
	"strings"
)

// FromMorseCode - Convert morse code to an ASCII string.
//
//	(c) 2023 Sam Caldwell.  MIT License
func FromMorseCode(morseCode string) (string, error) {
	morseCode = strings.TrimSpace(morseCode)
	morseWords := strings.Split(morseCode, "/")
	output := ""

	for wordPos, word := range morseWords {
		morseChars := strings.Split(word, " ")
		for charPos, morseChar := range morseChars {
			if asciiChar, ok := morseCodeToASCII[morseChar]; ok {
				ansi2.Green().Printf("%d:%d[%s] ", wordPos, charPos, asciiChar)
				output += asciiChar
			} else {
				if strings.TrimSpace(morseChar) != "" {
					return output, fmt.Errorf("invalid morseChar: '%s'", morseChar)
				}
			}
		}
		output += " "
	}
	ansi2.LF().Reset()

	return strings.TrimSpace(output), nil
}
