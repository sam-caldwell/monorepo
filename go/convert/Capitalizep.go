package convert

import "unicode"

func Capitalizep(s *string) string {
	if *s == "" {
		return *s
	}
	runes := []rune(*s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
