package terminal

import "github.com/sam-caldwell/monorepo/go/ansi"

const (
	Red uint8 = iota
	Green
	Yellow
)

func BracketBox(text string, textColor uint8, highlight bool) *ansi.Color {
	ansi.White().Dim().Printf("[")
	switch textColor {
	case Red:
		ansi.Red()
	case Green:
		ansi.Green()
	case Yellow:
		ansi.Yellow()
	default:
		ansi.White()
	}
	if highlight {
		ansi.BgYellow().Blink()
	} else {
		ansi.BgBlack()
	}
	return ansi.Bold().Print(text).Reset().
		White().Dim().Printf("]")
}
